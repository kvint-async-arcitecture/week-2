package password

import (
	"gitlab.com/go-ledger/common/errors"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrEncrypt = errors.InternalError("errEncrypt")
	ErrDecrypt = errors.InternalError("errDecrypt")
)

func Encrypt(pwd string) (string, error) {
	pwdHash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		return "", ErrEncrypt.Wrap(err)
	}

	return string(pwdHash), nil
}

func CheckPasswordHash(pwd, hash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pwd))
	if err != nil {
		return ErrDecrypt.Wrap(err)
	}

	return nil
}
