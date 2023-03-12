package user

type UUIDGenerator interface {
	Generate() string
}

type PasswordEncryptor interface {
	Encrypt(password string) string
}

type Repository interface {
	Create(u *User) error
	Get(id string) (*User, error)
	Update(u *User) error
	Delete(u *User) error
}
