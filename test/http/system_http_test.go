package http_test

import (
	"context"
	"encoding/json"
	"fmt"
	"gin_vue_admin_framework/common"
	"gin_vue_admin_framework/initialize"
	"gin_vue_admin_framework/internal/routes"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"time"
)

type tokenResponseData struct {
	Token string `json:"token"`
}
type LoginResponse struct {
	Data   tokenResponseData `json:"data"`
	Status string            `json:"status"`
	Msg    string            `json:"msg"`
}

var r *gin.Engine

func init() {
	initialize.InitSystem()

	r = routes.Router
}

func TestSysLogin(t *testing.T) {
	fmt.Println("mode", gin.Mode())
	postData := url.Values{
		"username":   {"admin"},
		"password":   {"123456"},
		"captcha":    {"123456"},
		"captcha_id": {"1"},
		"auto_login": {"null"},
	}

	reqBody := strings.NewReader(postData.Encode())
	req, err := http.NewRequest("POST", "/public/admin/user/login", reqBody)
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

	println("resBody:", resBodyString)
	var parseResp LoginResponse
	if err := json.Unmarshal(resBody, &parseResp); err != nil {
		t.Fatal("json.Unmarshal error:", err)
	}
	rdb := common.COM_REDIS
	if rdb == nil {
		t.Error("rdb is nil")
	}
	ctx := context.Background()
	_, err = rdb.Set(ctx, "x-header-token", parseResp.Data.Token, time.Duration(7)*24*time.Hour).Result()
	if err != nil {
		t.Error(err)
	}
	defer res.Body.Close()
}

func TestSysUpdateSysSettings(t *testing.T) {
	postData := make([]map[string]string, 0)
	data1 := make(map[string]string)
	data1["set_key"] = "sys_name"
	data1["set_value"] = "gvaf"
	postData = append(postData, data1)

	jsonData, _ := json.Marshal(postData)
	reqBody := strings.NewReader(string(jsonData))
	fmt.Println("reqBody:", reqBody)
	req, err := http.NewRequest("POST", "/sys/update_sys_settings", reqBody)
	if err != nil {
		t.Fatal("http.NewRequest error:", err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	rdb := common.COM_REDIS
	if rdb == nil {
		t.Error("rdb is nil")
	}
	ctx := context.Background()
	token, err := rdb.Get(ctx, "x-header-token").Result()
	if err != nil {
		t.Error(err)
	}
	req.Header.Set("x-header-token", token)

	rec := httptest.NewRecorder()

	r.ServeHTTP(rec, req)
	res := rec.Result()

	defer res.Body.Close()

	resBody, _ := io.ReadAll(res.Body)
	resBodyString := string(resBody)
	if res.StatusCode != http.StatusOK {
		t.Fatalf("status code error,want %d,got %d, msg: %v", http.StatusOK, res.StatusCode, resBodyString)
	}

}
