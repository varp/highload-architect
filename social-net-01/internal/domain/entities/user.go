package entities

type User struct {
	Id         string
	Age        int
	Biography  string
	City       string
	FirstName  string
	SecondName string
	Password   string
}

func (u *User) SetPassword(newPassword string) error {
	return nil
}

type UserRepository interface {
	Get(id string) (*User, error)
	Save(u *User) error
}
