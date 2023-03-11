package user

import (
	"go.vardan.dev/highload-architect/social-net-01/internal/domain/models/user"
)

type Usecase interface {
	Login(userId, password string) error
	Register(u *user.User) error
	Get(userId string) (*user.User, error)
}

type userUsecase struct {
	ur user.Repository
}

func New(ur user.Repository) Usecase {
	return &userUsecase{ur: ur}
}

func (u *userUsecase) Login(userId, password string) error {
	// TODO implement me
	panic("implement me")
}

func (u *userUsecase) Register(user *user.User) error {
	// TODO implement me
	panic("implement me")
}

func (u *userUsecase) Get(userId string) (*user.User, error) {
	// TODO implement me
	panic("implement me")
}
