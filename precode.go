package main

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var cafeList = map[string][]string{
	"moscow": {"Мир кофе", "Сладкоежка", "Кофе и завтраки", "Сытый студент"},
}

func mainHandle(w http.ResponseWriter, req *http.Request) {
	countStr := req.URL.Query().Get("count")
	if countStr == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("count missing"))
		return
	}

	count, err := strconv.Atoi(countStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("wrong count value"))
		return
	}

	city := req.URL.Query().Get("city")

	cafe, ok := cafeList[city]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("wrong city value"))
		return
	}

	if count > len(cafe) {
		count = len(cafe)
	}

	answer := strings.Join(cafe[:count], ",")

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(answer))
}

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	// Общее количество кафе в городе "Москва" - 4
	totalCount := 4
	city := "moscow"
	count := totalCount + 1 // Запрашиваем больше, чем доступно кафе

	req, err := http.NewRequest("GET", "/?city="+city+"&count="+strconv.Itoa(count), nil)
	if err != nil {
		t.Fatal(err)
	}

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	// Проверяем, что статус ответа - OK
	assert.Equal(t, http.StatusOK, responseRecorder.Code)

	// Извлекаем тело ответа
	body := responseRecorder.Body.String()

	// Проверяем, что тело ответа содержит все кафе, разделенные запятыми
	expectedBody := strings.Join(cafeList[city], ",")
	assert.Equal(t, expectedBody, body)
}
