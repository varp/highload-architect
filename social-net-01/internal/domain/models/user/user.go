package user

import "fmt"

type User struct {
	Id           string
	Age          int
	Biography    string
	City         string
	FirstName    string
	SecondName   string
	PasswordHash string
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

func New(age int, biography, city, firstName, secondName, password string) *User {
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

	encrypted := passwordEncryptor.Encrypt(password)
	return u.PasswordHash == encrypted
}
