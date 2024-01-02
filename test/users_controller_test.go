package test

import (
	"amikom-pedia-api/app"
	"amikom-pedia-api/controller"
	"amikom-pedia-api/helper"
	"amikom-pedia-api/repository"
	"amikom-pedia-api/service"
	"database/sql"
	"encoding/json"
	"github.com/go-playground/validator/v10"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func setupTestDB() *sql.DB {
	db, err := sql.Open("postgres", "postgres://root:root@localhost:5432/amikom_pedia?sslmode=disable")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}

func setupRouter(db *sql.DB) http.Handler {
	validate := validator.New()
	categoryRepository := repository.NewUserRepository()
	categoryService := service.NewUserService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)
	router := app.NewRouter(categoryController)

	return router
}

func truncateCategory(db *sql.DB) {
	db.Exec("TRUNCATE users")
}

func TestCreateCategorySuccess(t *testing.T) {
	db := setupTestDB()
	truncateCategory(db)
	router := setupRouter(db)
	requestJson := `{ "id": "20.11.3785", "username": "bene", "displayname": "testing bang", "email": "bene@students.amikom.ac.id", "password": "waw"}`
	requestBody := strings.NewReader(requestJson)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/users", requestBody)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
	assert.Equal(t, "20.11.3785", responseBody["data"].(map[string]interface{})["id"])
}
