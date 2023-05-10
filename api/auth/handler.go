package auth

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// 以下にユーザー認証用のハンドラーを実装
type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var user = map[string]string{
	"user1": "password1",
}

var jwtKey = []byte("your_secret_key")

func Login() {
	
}
