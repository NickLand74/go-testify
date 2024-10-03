package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	req := httptest.NewRequest("GET", "/?city=moscow&count=10", nil)
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)

	handler.ServeHTTP(responseRecorder, req)

	if status := responseRecorder.Code; status != http.StatusOK {
		t.Errorf("handler returned неверный код состояния: got %v want %v", status, http.StatusOK)
	}

	expected := "Мир кофе,Сладкоежка,Кофе и завтраки,Сытый студент"
	if responseRecorder.Body.String() != expected {
		t.Errorf("handler вернул неожидаемое тело: got %v want %v", responseRecorder.Body.String(), expected)
	}
	// Этот тест проверяет, правильно ли хендлер возвращает все доступные кафе, когда передано значение `count`, превышающее общее количество кафе
}

func TestMainHandlerWhenCityIsWrong(t *testing.T) {
	req := httptest.NewRequest("GET", "/?city=invalidcity&count=2", nil)
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)

	handler.ServeHTTP(responseRecorder, req)

	if status := responseRecorder.Code; status != http.StatusBadRequest {
		t.Errorf("handler вернул неверный код состояния: got %v want %v", status, http.StatusBadRequest)
	}

	expected := "wrong city value"
	if responseRecorder.Body.String() != expected {
		t.Errorf("handler вернул неожиданное тело: got %v want %v", responseRecorder.Body.String(), expected)
	}
	// Этот тест проверяет, отвечает ли хендлер кодом `400 Bad Request` и соответствующим сообщением об ошибке, если передан неправильный город
}

func TestMainHandlerWhenCountIsMissing(t *testing.T) {
	req := httptest.NewRequest("GET", "/?city=moscow", nil)
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)

	handler.ServeHTTP(responseRecorder, req)

	if status := responseRecorder.Code; status != http.StatusBadRequest {
		t.Errorf("handler вернул неверный код состояния: got %v want %v", status, http.StatusBadRequest)
	}

	expected := "count missing"
	if responseRecorder.Body.String() != expected {
		t.Errorf("handler вернул неожиданное тело: got %v want %v", responseRecorder.Body.String(), expected)
	}
	// Этот тест проверяет, отвечает ли хендлер кодом `400 Bad Request` и соответствующим сообщением об ошибке, когда параметр `count` отсутствует в запросе.
	// действительно файл precode_test.go почему то не подгрузился в прошлый раз
}
