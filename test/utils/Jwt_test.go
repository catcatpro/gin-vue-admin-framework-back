package utils

import (
	"gin_vue_admin_framework/utils"
	"testing"
)

func TestJWT(t *testing.T) {
	j := utils.NewJWT()
	claims := j.CreateClaims(1, "user1")
	token, err := j.CreateToken(claims)
	if err != nil {
		t.Error(err)
	}
	claims2, err := j.ParseToken(token)
	if err != nil {
		t.Error(err)
	}
	if claims2.Username != "user1" {
		t.Error("should be user1 for parse token")
	}
	//t.Error("test")
}
