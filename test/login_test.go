package test

import (
	"bytes"
	"database/sql"
	autoriz "diplom/src/autorization"
	models "diplom/src/database"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	_ "github.com/lib/pq"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestSignupHandler(t *testing.T) {
	// Создаем тестовый запрос
	user := models.User{
		Username: "testuser",
		Password: "testpassword",
	}
	db, err := sql.Open("postgres", "user=postgres password=postgres dbname=diplom_rob sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	payload, _ := json.Marshal(user)
	request, _ := http.NewRequest("POST", "/signup", bytes.NewBuffer(payload))

	// Создаем тестовый контекст и рекордер ответов
	recorder := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(recorder)
	context.Request = request

	// Вызываем функцию SignupHandler
	autoriz.SignupHandler(context, db)

	// Проверяем код состояния и ожидаемый ответ
	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Contains(t, recorder.Body.String(), "\"username\":\"testuser\"")
}

func TestLoginHandler(t *testing.T) {
	// Создаем тестовый запрос
	creds := models.Credentials{
		Username: "testuser",
		Password: "testpassword",
	}
	db, err := sql.Open("postgres", "user=postgres password=postgres dbname=diplom_rob sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	payload, _ := json.Marshal(creds)
	request, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(payload))

	// Создаем тестовый контекст и рекордер ответов
	recorder := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(recorder)
	context.Request = request

	// Вызываем функцию LoginHandler
	autoriz.LoginHandler(context, db)

	// Проверяем код состояния и ожидаемый ответ
	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Contains(t, recorder.Body.String(), "Login successful")
}
