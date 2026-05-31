package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

// hmacSampleSecret is the shared HMAC signing key used for JWT generation
// and validation in this project.
//
// Note: This key is hard-coded for simplicity in the current architecture.
// In production systems, it should be loaded from secure configuration.
var hmacSampleSecret = []byte("1231")

// BCYHashPassword hashes the plaintext password in place using bcrypt.
//
// Input:
//   - pass: pointer to plaintext password string
//
// Output:
//   - the value pointed to by pass is replaced with its bcrypt hash.
//
// The function intentionally ignores the hash error to keep compatibility with
// the existing code flow.
func BCYHashPassword(pass *string) {
	bytePass := []byte(*pass)
	hPass, _ := bcrypt.GenerateFromPassword(bytePass, bcrypt.DefaultCost)
	*pass = string(hPass)
}

// BCYComparePassword compares a stored bcrypt hash with a plaintext password.
//
// Parameters:
//   - dbPass: hashed password stored in database/config
//   - pass: plaintext password provided by user
//
// Returns true only when the plaintext password matches the hash.
func BCYComparePassword(dbPass, pass string) bool {
	return bcrypt.CompareHashAndPassword([]byte(dbPass), []byte(pass)) == nil
}

// BCYGenerateToken creates a signed JWT token for the given user.
//
// Claims:
//   - exp: expiration timestamp (current time + 1 hour)
//   - iat: issued-at timestamp
//   - user: username
//
// Returns:
//   - token string
//   - expiration unix timestamp (seconds)
//
// The function uses HS256 and the shared hmacSampleSecret key.
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

// BCYValidateToken parses and validates the given JWT token string.
//
// Validation behavior:
//   - Ensures the token's signing algorithm is HMAC-based.
//   - Uses hmacSampleSecret as verification key.
//
// Returns the parsed token and any validation/parsing error.
func BCYValidateToken(token string) (*jwt.Token, error) {
	// The callback provides the verification key after checking the signing method.
	return jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			// Reject tokens signed with unsupported algorithms.
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return hmacSampleSecret, nil
	})
}
