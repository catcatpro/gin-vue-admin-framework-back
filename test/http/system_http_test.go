package http_test

import (
	"gin_vue_admin_framework/initialize"
	"gin_vue_admin_framework/internal/routes"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestSysLogin(t *testing.T) {
	initialize.InitSystem()

	r := routes.Router
	postData := url.Values{
		"username":   {"admin"},
		"password":   {"123456"},
		"captcha":    {"123456"},
		"captcha_id": {"1"},
	}

	reqBody := strings.NewReader(postData.Encode())
	req, err := http.NewRequest("POST", "/public/sys/login", reqBody)
	if err != nil {
		t.Fatal("http.NewRequest error:", err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	rec := httptest.NewRecorder()

	r.ServeHTTP(rec, req)

	res := rec.Result()
	resBody, _ := io.ReadAll(res.Body)
	resBodyString := string(resBody)
	if res.StatusCode != http.StatusOK {
		t.Fatalf("status code error,want %d,got %d, msg: %v", http.StatusOK, res.StatusCode, resBodyString)
	}

	defer res.Body.Close()
	
}
