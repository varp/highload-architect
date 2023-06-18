package tools

import (
	"go.vardan.dev/highload-architect/social-net-01/internal/domain/models"
	"golang.org/x/crypto/bcrypt"
)

type PasswordEncryptor struct{}

func NewPasswordEncryptor() models.PasswordEncryptor {
	return &PasswordEncryptor{}
}

func (p *PasswordEncryptor) Encrypt(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		panic(err)
	}
	return string(bytes)
}

func (p *PasswordEncryptor) Compare(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
