package Session

import (
	"encoding/json"
	"gin_vue_admin_framework/internal/routes"
	"github.com/go-playground/assert/v2"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

type session struct {
	Test string `json:"test"`
}

func TestSessionCookie(t *testing.T) {
	var w *httptest.ResponseRecorder
	r := routes.Router
	url := "/example/test_session"
	req, _ := http.NewRequest("GET", url, nil)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	m := session{}
	body, _ := io.ReadAll(w.Body)
	err := json.Unmarshal(body, &m)
	if err != nil {
		log.Fatal(err)
	}
	assert.Equal(t, m.Test, "test_value")
}
