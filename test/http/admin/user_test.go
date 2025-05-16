package admin

import (
	"gin_vue_admin_framework/initialize"
	"gin_vue_admin_framework/internal/routes"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

var r *gin.Engine

func init() {
	initialize.InitSystem()

	r = routes.Router
}

func TestCreateUser(t *testing.T) {
	postData := url.Values{
		"username":         {"user1"},
		"password":         {"123456"},
		"confirm_password": {"123456"},
	}

	reqBody := strings.NewReader(postData.Encode())
	req, err := http.NewRequest("POST", "/admin/user/create", reqBody)
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
