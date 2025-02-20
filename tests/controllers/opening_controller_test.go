package controllers_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

// func TestMain(m *testing.M) {

// }

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("{}"))
}

func TestShowOpening(t *testing.T) {
	req := httptest.NewRequest("GET", "/opening?id=2", nil)
	w := httptest.NewRecorder()

	handler(w, req)

	resp := w.Result()

	fmt.Println(resp.StatusCode)

	assert.Equal(t, resp.StatusCode, http.StatusOK, "Show opening should return status 200")

}

func TestShowOpenings(t *testing.T) {
	req := httptest.NewRequest("GET", "/openings", nil)
	w := httptest.NewRecorder()

	handler(w, req)

	resp := w.Result()

	fmt.Println(resp.StatusCode)

	assert.Equal(t, resp.StatusCode, http.StatusOK, "Show openings should return status 200")

}
