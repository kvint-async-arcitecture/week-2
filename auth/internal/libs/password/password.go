package password

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

var (
	ErrEncrypt = errors.New("errEncrypt")
	ErrDecrypt = errors.New("errDecrypt")
)

func Encrypt(pwd string) (string, error) {
	pwdHash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		return "", ErrEncrypt
	}

	return string(pwdHash), nil
}

func CheckPasswordHash(pwd, hash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pwd))
	if err != nil {
		return ErrDecrypt
	}

	return nil
}
