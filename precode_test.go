package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func performRequest(t *testing.T, method, path string) *httptest.ResponseRecorder {
	req, err := http.NewRequest(method, path, nil)
	assert.NoError(t, err)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	return responseRecorder
}

func TestMainHandlerWhenCountEqualsTotal(t *testing.T) {
	responseRecorder := performRequest(t, "GET", "/cafe?count=4&city=moscow")

	assert.Equal(t, http.StatusOK, responseRecorder.Code)
	assert.NotEmpty(t, responseRecorder.Body.String())
}

func TestMainHandlerWithUnknownCity(t *testing.T) {
	responseRecorder := performRequest(t, "GET", "/cafe?count=4&city=nonexistent")

	assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)
	assert.Contains(t, responseRecorder.Body.String(), "wrong city value")
}

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	responseRecorder := performRequest(t, "GET", "/cafe?count=10&city=moscow")

	assert.Equal(t, http.StatusOK, responseRecorder.Code)
	assert.NotEmpty(t, responseRecorder.Body.String())
}
