package utils

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"time"
)

//密钥
var hmacSampleSecret = []byte("1231")

// BCYHashPassword 生成密码
func BCYHashPassword(pass *string) {
	bytePass := []byte(*pass)
	hPass, _ := bcrypt.GenerateFromPassword(bytePass, bcrypt.DefaultCost)
	*pass = string(hPass)
}

// BCYComparePassword 密码比对
func BCYComparePassword(dbPass, pass string) bool {
	return bcrypt.CompareHashAndPassword([]byte(dbPass), []byte(pass)) == nil
}

// BCYGenerateToken -> generates token 创建token
func BCYGenerateToken(user string) (string, int64) {
	expInt := time.Duration(3600) * time.Second
	exp := time.Now().Add(expInt).Unix()
	claims := jwt.MapClaims{
		"exp":  exp,
		"iat":  time.Now().Unix(),
		"user": user,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, _ := token.SignedString(hmacSampleSecret)
	return t, exp
}

// BCYValidateToken --> validate the given token
func BCYValidateToken(token string) (*jwt.Token, error) {
	//2nd arg function return secret key after checking if the signing method is HMAC and returned key is used by 'Parse' to decode the token)
	return jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			//nil secret key
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return hmacSampleSecret, nil
	})
}
