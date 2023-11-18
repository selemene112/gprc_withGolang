package Middleware

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

// HashPassword menghasilkan hash dari kata sandi
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// ComparePasswords membandingkan hash kata sandi dengan kata sandi yang diberikan
func ComparePasswords(hashedPassword, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		// Verifikasi kata sandi gagal
		return errors.New("password salah")
	}
	return nil
}
