package models

import (
	"fmt"
)

type UUIDGenerator interface {
	Generate() string
}

type PasswordEncryptor interface {
	Encrypt(password string) string
	Compare(hashedPassword, password string) bool
}

type UserRepository interface {
	Create(u *User) error
	Get(id string) (*User, error)
	Update(u *User) error
	Delete(u *User) error
	Search(firstName, lastName string) ([]*User, error)
}

type User struct {
	Id           string `db:"id"`
	Age          int    `db:"age"`
	Biography    string `db:"biography"`
	City         string `db:"city"`
	FirstName    string `db:"first_name"`
	SecondName   string `db:"second_name"`
	PasswordHash string `db:"password_hash"`
}

var (
	idGenerator       UUIDGenerator
	passwordEncryptor PasswordEncryptor
)

func Configure(generator UUIDGenerator, encryptor PasswordEncryptor) {
	idGenerator = generator
	passwordEncryptor = encryptor
}

func checkConfiguration() {
	if idGenerator == nil {
		panic(fmt.Errorf("confgiuration error: idGenerator is nil"))
	}

	if passwordEncryptor == nil {
		panic(fmt.Errorf("configuration error: passwordEncryptor is nil"))
	}
}

func NewUser(age int, biography, city, firstName, secondName, password string) *User {
	checkConfiguration()

	u := User{
		Id:           idGenerator.Generate(),
		Age:          age,
		Biography:    biography,
		City:         city,
		FirstName:    firstName,
		SecondName:   secondName,
		PasswordHash: passwordEncryptor.Encrypt(password),
	}

	return &u
}

func (u *User) SetPassword(password string) {
	checkConfiguration()

	u.PasswordHash = passwordEncryptor.Encrypt(password)
}

func (u *User) CheckPassword(password string) bool {
	checkConfiguration()

	return passwordEncryptor.Compare(u.PasswordHash, password)
}
