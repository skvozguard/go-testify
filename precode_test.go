package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4
	req := httptest.NewRequest("GET", fmt.Sprintf("/cafe?count=%d&city=moscow", totalCount+1), nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusOK, responseRecorder.Code)
	bodyString, err := io.ReadAll(responseRecorder.Body)
	assert.NoError(t, err)
	assert.NotEmpty(t, bodyString)

	cafes := strings.Split(string(bodyString), ",")
	assert.Equal(t, totalCount, len(cafes))
}
func TestMainHandlerCorrectRequest(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=2&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, http.StatusOK, responseRecorder.Code)
	bodyBytes, err := io.ReadAll(responseRecorder.Body)
	require.NoError(t, err)
	require.NotEmpty(t, bodyBytes)
}

func TestMainHandlerUnsupportedCity(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=2&city=unsupportedCity", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)
	bodyString := responseRecorder.Body.String()
	assert.Contains(t, bodyString, "wrong city value")
}
