package utils

import "testing"

func TestBCYHashPassword(t *testing.T) {
	pwd := "123456"
	t.Log("OldPwd:", pwd)
	BCYHashPassword(&pwd)
	//$2a$10$WqH8Su/nNyzUHyd258iK7e3g2.f3u4a57KTleSDzQwil7YoLq9kPy
	t.Log("NewPwd:", pwd)
}

func TestBCYComparePassword(t *testing.T) {
	pwd := "123456"
	keyPwd := `$2a$10$VTiL7jOvrdShutUePqyPJ./Yjc72TLarQM60cvnQ3PA/INx/jyb2a`
	if BCYComparePassword(pwd, keyPwd) {
		t.Log("pwd check success")
	} else {
		t.Error("BCYComparePassword failed")
	}
}

func TestBCYGenerateToken(t *testing.T) {
	user := "admin"
	token, exp := BCYGenerateToken(user)
	t.Log("BCYGenerateToken:", token, " ;exp:", exp)
}

func TestBCYValidateToken(t *testing.T) {

}
