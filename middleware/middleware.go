package middleware

import (
	"amikom-pedia-api/helper"
	"amikom-pedia-api/model/web"
	"amikom-pedia-api/utils"
	"amikom-pedia-api/utils/token"
	"errors"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"strings"
	"time"
)

const (
	authorizationHeaderKey  = "authorization"
	authorizationTypeBearer = "bearer"
	AuthorizationPayloadKey = "authorization_payload"
)

type Middleware struct {
	Handler    http.Handler
	TokenMaker token.Maker
}

func NewMiddleware(handler http.Handler, tokenMaker token.Maker) *Middleware {
	return &Middleware{Handler: handler, TokenMaker: tokenMaker}
}

func (m *Middleware) AuthMiddleware(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		authorizationHeader := r.Header.Get(authorizationHeaderKey)
		fields := strings.Fields(authorizationHeader)

		if len(authorizationHeaderKey) == 0 {

			err := errors.New("authorization header is not provided")
			m.unauthorizedResponse(w, err.Error())

		} else if len(fields) < 2 {

			err := errors.New("invalid authorization header format")
			m.unauthorizedResponse(w, err.Error())

		} else if strings.ToLower(fields[0]) != authorizationTypeBearer { // bearer token is the only supported authorization type
			err := errors.New("authorization type is not supported")
			m.unauthorizedResponse(w, err.Error())
		} else {
			accessToken := fields[1]
			payload, err := m.TokenMaker.VerifyToken(accessToken)

			if err != nil {
				m.unauthorizedResponse(w, err.Error())
			} else {
				r.Header.Set(AuthorizationPayloadKey, utils.FromUsernameAndUUIDToString(payload.Username, payload.UserID))
				//m.Handler.ServeHTTP(w, r)
				next(w, r, params)
			}
		}

	}
}

func (m *Middleware) LoggingMiddleware(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		// Record the start time
		startTime := time.Now()

		// Log information about the incoming request
		log.Printf("[%s] %s %s", r.Method, r.URL.Path, r.RemoteAddr)

		// Call the next handler
		next(w, r, params)

		// Record the end time
		endTime := time.Now()

		// Calculate and log the request duration
		duration := endTime.Sub(startTime)
		log.Printf("Request duration: %v", duration)
	}
}

func (m *Middleware) WrapperMiddleware(next httprouter.Handle) httprouter.Handle {
	return m.LoggingMiddleware(m.AuthMiddleware(next))
}

func (m *Middleware) unauthorizedResponse(writer http.ResponseWriter, error string) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusUnauthorized)

	webResponse := web.WebResponse{
		Code:   http.StatusUnauthorized,
		Status: error,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
