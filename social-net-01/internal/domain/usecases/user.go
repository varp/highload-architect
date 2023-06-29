package usecases

import (
	"fmt"
	"go.vardan.dev/highload-architect/social-net-01/internal/domain/models"
)

type UserUsecase interface {
	Login(userId, password string) error
	Register(u *models.User) error
	Get(userId string) (*models.User, error)
	Search(firstName, lastName string) ([]*models.User, error)
}

type userUsecase struct {
	ur models.UserRepository
	ug models.UUIDGenerator
	pe models.PasswordEncryptor
}

func NewUserUsecase(ur models.UserRepository, ug models.UUIDGenerator, pe models.PasswordEncryptor) UserUsecase {
	models.Configure(ug, pe)
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

func (uu *userUsecase) Register(user *models.User) error {
	return uu.ur.Create(user)
}

func (uu *userUsecase) Get(userId string) (*models.User, error) {
	return uu.ur.Get(userId)
}

func (uu *userUsecase) Search(firstName, lastName string) ([]*models.User, error) {
	return uu.ur.Search(firstName, lastName)
}
