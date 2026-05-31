package utils

import "testing"

// TestBCYHashPassword verifies that plaintext can be transformed into a bcrypt hash.
// The test prints before/after values for manual inspection.
func TestBCYHashPassword(t *testing.T) {
	pwd := "123456"
	t.Log("OldPwd:", pwd)
	BCYHashPassword(&pwd)
	// Example hash format: $2a$10$...
	t.Log("NewPwd:", pwd)
}

// TestBCYComparePassword verifies that known plaintext matches a known bcrypt hash.
func TestBCYComparePassword(t *testing.T) {
	pwd := "123456"
	keyPwd := `$2a$10$VTiL7jOvrdShutUePqyPJ./Yjc72TLarQM60cvnQ3PA/INx/jyb2a`
	if BCYComparePassword(keyPwd, pwd) {
		t.Log("pwd check success")
	} else {
		t.Error("BCYComparePassword failed")
	}
}

// TestBCYGenerateToken ensures token generation returns a token string and expiration.
func TestBCYGenerateToken(t *testing.T) {
	user := "admin"
	token, exp := BCYGenerateToken(user)
	t.Log("BCYGenerateToken:", token, " ;exp:", exp)
}

// TestBCYValidateToken is currently reserved for future token-parse assertions.
func TestBCYValidateToken(t *testing.T) {

}
