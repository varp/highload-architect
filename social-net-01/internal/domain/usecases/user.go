package usecases

import "go.vardan.dev/highload-architect/social-net-01/internal/domain/entities"

type UserUsecase interface {
	Login(userId, password string) error
	Register(u *entities.User) error
	Get(userId string) (*entities.User, error)
}

type userUsecase struct {
	ur entities.UserRepository
}

func NewUserUsecase(ur entities.UserRepository) UserUsecase {
	return &userUsecase{ur: ur}
}

func (u *userUsecase) Login(userId, password string) error {
	// TODO implement me
	panic("implement me")
}

func (u *userUsecase) Register(user *entities.User) error {
	// TODO implement me
	panic("implement me")
}

func (u *userUsecase) Get(userId string) (*entities.User, error) {
	// TODO implement me
	panic("implement me")
}
