package models

type UUIDGenerator interface {
	Generate() string
}

type PasswordEncryptor interface {
	Encrypt(password string) string
	Compare(hashedPassword, password string) bool
}

type UserRepository interface {
	Create(u *User) (int64, error)
	Get(id string) (*User, error)
	Update(u *User) error
	Delete(u *User) error
	Search(firstName, lastName string) ([]*User, error)
}

type User struct {
	Id           int64  `db:"id"`
	Age          int    `db:"age"`
	Biography    string `db:"biography"`
	City         string `db:"city"`
	FirstName    string `db:"first_name"`
	SecondName   string `db:"second_name"`
	PasswordHash string `db:"password_hash"`
}

func (u *User) SetPassword(passwordEncryptor PasswordEncryptor, password string) {
	u.PasswordHash = passwordEncryptor.Encrypt(password)
}

func (u *User) CheckPassword(passwordEncryptor PasswordEncryptor, password string) bool {
	return passwordEncryptor.Compare(u.PasswordHash, password)
}
