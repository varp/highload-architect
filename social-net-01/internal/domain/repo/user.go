package repo

import (
	"go.vardan.dev/highload-architect/social-net-01/internal/domain/entities"
)

type userRepository struct {
}

func NewUserRepository() entities.UserRepository {
	return &userRepository{}
}

func (u *userRepository) Get(id string) (*entities.User, error) {
	// TODO implement me
	panic("implement me")
}

func (u *userRepository) Save(user *entities.User) error {
	// TODO implement me
	panic("implement me")
}
