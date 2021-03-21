package main

import (
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func setupTestEnv() *Env {
	env := initDatabase()

	return env

}

func TestDeleteNegativeIDEndpoint(t *testing.T) {
	req := httptest.NewRequest("DELETE", "http://localhost:8080/api/v1/movies/delete/-1", nil)

	// http.Response
	resp := httptest.NewRecorder()

	router := setupRoutes(setupTestEnv())
	router.ServeHTTP(resp, req)

	assert.Equal(t, 404, resp.Result().StatusCode)

}

func TestGetAllMovies(t *testing.T) {
	req := httptest.NewRequest("GET", "http://localhost:8080/api/v1/movies", nil)

	// http.Response
	resp := httptest.NewRecorder()

	router := setupRoutes(setupTestEnv())
	router.ServeHTTP(resp, req)

	assert.Equal(t, 500, resp.Result().StatusCode)

}
