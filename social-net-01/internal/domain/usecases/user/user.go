package user

import (
	"fmt"

	"go.vardan.dev/highload-architect/social-net-01/internal/domain/models/user"
)

type Usecase interface {
	Login(userId, password string) error
	Register(u *user.User) error
	Get(userId string) (*user.User, error)
}

type userUsecase struct {
	ur user.Repository
	ug user.UUIDGenerator
	pe user.PasswordEncryptor
}

func New(ur user.Repository, ug user.UUIDGenerator, pe user.PasswordEncryptor) Usecase {
	user.Configure(ug, pe)
	return &userUsecase{ur, ug, pe}
}

func (uu *userUsecase) Login(userId, password string) error {
	retErr := fmt.Errorf("invalid userId or password was supplied")

	u, err := uu.ur.Get(userId)
	if err != nil {
		return retErr
	}

	if !u.CheckPassword(password) {
		return retErr
	}

	return nil
}

func (uu *userUsecase) Register(user *user.User) error {
	err := uu.ur.Create(user)
	if err != nil {
		return err
	}

	return nil
}

func (uu *userUsecase) Get(userId string) (*user.User, error) {
	u, err := uu.ur.Get(userId)
	if err != nil {
		return nil, err
	}

	return u, nil
}
