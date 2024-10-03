package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMainHandlerWithValidRequest(t *testing.T) {
	req := httptest.NewRequest("GET", "/?city=moscow&count=2", nil)
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)

	handler.ServeHTTP(responseRecorder, req)

	// Проверяем код состояния
	require.Equal(t, http.StatusOK, responseRecorder.Code, "handler returned неверный код состояния")

	// Проверяем, что тело ответа не пустое
	assert.NotEmpty(t, responseRecorder.Body.String(), "тело ответа должно быть не пустым")
}

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	req := httptest.NewRequest("GET", "/?city=moscow&count=10", nil)
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)

	handler.ServeHTTP(responseRecorder, req)

	// Проверяем код состояния
	require.Equal(t, http.StatusOK, responseRecorder.Code, "handler returned неверный код состояния")

	expected := "Мир кофе,Сладкоежка,Кофе и завтраки,Сытый студент"
	// Проверяем тело ответа
	require.Equal(t, expected, responseRecorder.Body.String(), "handler вернул неожидаемое тело")

	// Этот тест проверяет, правильно ли хендлер возвращает все доступные кафе, когда передано значение count, превышающее общее количество кафе
}

func TestMainHandlerWhenCityIsWrong(t *testing.T) {
	req := httptest.NewRequest("GET", "/?city=invalidcity&count=2", nil)
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)

	handler.ServeHTTP(responseRecorder, req)

	// Проверяем код состояния
	require.Equal(t, http.StatusBadRequest, responseRecorder.Code, "handler вернул неверный код состояния")

	expected := "wrong city value"
	// Проверяем тело ответа
	require.Equal(t, expected, responseRecorder.Body.String(), "handler вернул неожиданное тело")

	// Этот тест проверяет, отвечает ли хендлер кодом 400 Bad Request и соответствующим сообщением об ошибке, если передан неправильный город
}

func TestMainHandlerWhenCountIsMissing(t *testing.T) {
	req := httptest.NewRequest("GET", "/?city=moscow", nil)
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)

	handler.ServeHTTP(responseRecorder, req)

	// Проверяем код состояния
	require.Equal(t, http.StatusBadRequest, responseRecorder.Code, "handler вернул неверный код состояния")

	expected := "count missing"
	// Проверяем тело ответа
	require.Equal(t, expected, responseRecorder.Body.String(), "handler вернул неожиданное тело")

	// Проверяем, что тело ответа не пустое
	assert.NotEmpty(t, responseRecorder.Body.String(), "тело ответа должно быть не пустым")

	// Этот тест проверяет, отвечает ли хендлер кодом 400 Bad Request и соответствующим сообщением об ошибке, когда параметр count отсутствует в запросе.
}
