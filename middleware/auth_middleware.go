package middleware

import (
	"amikom-pedia-api/helper"
	"amikom-pedia-api/model/web"
	"amikom-pedia-api/utils/token"
	"errors"
	"net/http"
	"strings"
)

const (
	authorizationHeaderKey  = "authorization"
	authorizationTypeBearer = "bearer"
	authorizationPayloadKey = "authorization_payload"
)

type AuthMiddleware struct {
	Handler    http.Handler
	TokenMaker token.Maker
}

func NewAuthMiddleware(handler http.Handler, tokenMaker token.Maker) *AuthMiddleware {
	return &AuthMiddleware{Handler: handler, TokenMaker: tokenMaker}
}

func (middleware *AuthMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	authorizationHeader := request.Header.Get(authorizationHeaderKey)
	fields := strings.Fields(authorizationHeader)

	if request.URL.Path == "/api/v1/login" || request.URL.Path == "/api/v1/register" || request.URL.Path == "/api/v1/users" || request.URL.Path == "/api/v1/forgot-password" || request.URL.Path == "/api/v1/set-new-password" || request.URL.Path == "/api/v1/otp/validate" || request.URL.Path == "/api/v1/otp/send" || request.URL.Path == "/api/v1/users/change-password" {

		middleware.Handler.ServeHTTP(writer, request)

	} else if len(authorizationHeaderKey) == 0 {

		err := errors.New("authorization header is not provided")
		middleware.unauthorizedResponse(writer, err.Error())

	} else if len(fields) < 2 {

		err := errors.New("invalid authorization header format")
		middleware.unauthorizedResponse(writer, err.Error())

	} else if strings.ToLower(fields[0]) != authorizationTypeBearer { // bearer token is the only supported authorization type
		err := errors.New("authorization type is not supported")
		middleware.unauthorizedResponse(writer, err.Error())

	} else {
		accessToken := fields[1]
		payload, err := middleware.TokenMaker.VerifyToken(accessToken)

		if err != nil {
			middleware.unauthorizedResponse(writer, err.Error())
		} else {
			request.Header.Set(authorizationPayloadKey, payload.Username)
			middleware.Handler.ServeHTTP(writer, request)
		}
	}
}

func (middleware *AuthMiddleware) unauthorizedResponse(writer http.ResponseWriter, error string) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusUnauthorized)

	webResponse := web.WebResponse{
		Code:   http.StatusUnauthorized,
		Status: error,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
