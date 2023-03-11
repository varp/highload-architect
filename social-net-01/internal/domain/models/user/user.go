package user

type UUIDGenerator interface {
	Generate() string
}

type PasswordEncryptor interface {
	Encrypt(password string) string
}

type PasswordDecryptor interface {
	Decrypt(encrypted string) string
}

type User struct {
	Id         string
	Age        int
	Biography  string
	City       string
	FirstName  string
	SecondName string
	Password   string
}

func (u *User) SetPassword(password string, encryptor PasswordEncryptor) error {
	return nil
}

func (u *User) CheckPassword(password string, decryptor PasswordDecryptor) (bool, error) {
	return true, nil
}

func (u *User) SetID(uuidGenerator UUIDGenerator) {
	u.Id = uuidGenerator.Generate()
}

type Repository interface {
	Create(u *User) error
	Get(id string) (*User, error)
	Update(u *User) error
	Delete(u *User) error
}
