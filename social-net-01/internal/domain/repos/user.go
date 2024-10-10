package repos

import (
	"fmt"
	"go.vardan.dev/highload-architect/social-net-01/internal/domain/models"
	"log"

	"github.com/jmoiron/sqlx"
	"go.vardan.dev/highload-architect/social-net-01/internal/domain/repos/mysql"
)

var ErrUserNotFound = fmt.Errorf("user not found")

type userRepository struct {
	db *sqlx.DB
}

func NewUserRepository() models.UserRepository {
	db := mysql.Connect()

	return &userRepository{
		db: db,
	}
}

func (r *userRepository) Create(user *models.User) (userId int64, err error) {

	sql := "INSERT INTO users (age, biography, city, first_name, second_name, password_hash) VALUES (?, ?, ?, ?, ?, ?)"
	res, err := r.db.Exec(sql, user.Age, user.Biography, user.City, user.FirstName, user.SecondName, user.PasswordHash)

	if err != nil {
		return
	}

	userId, err = res.LastInsertId()
	if err != nil {
		return
	}

	return
}

func (r *userRepository) Get(id string) (*models.User, error) {
	row, err := r.db.Queryx("SELECT * FROM users WHERE id = ?", id)
	if err != nil {
		log.Fatalf("Failed to retrieve user with ID (%s) from DB", id)
		return nil, err
	}

	if !row.Next() {
		return nil, ErrUserNotFound
	}

	var u models.User
	err = row.StructScan(&u)
	if err != nil {
		log.Fatalf("Failed to map DB response to domain User model")
		return nil, err
	}

	return &u, nil
}

func (r *userRepository) Update(user *models.User) error {
	sql := "UPDATE users SET age = ?, biography = ?, city = ?, first_name = ?, second_name = ?, password_hash = ? WHERE id = ?"
	_, err := r.db.Exec(sql, user.Age, user.Biography, user.City, user.FirstName, user.SecondName,
		user.PasswordHash, user.Id)
	if err != nil {
		return err
	}

	return nil
}

func (r *userRepository) Delete(user *models.User) error {
	_, err := r.db.Exec("DELETE FROM users WHERE id = ?", user.Id)
	if err != nil {
		return err
	}

	return nil
}

func (r *userRepository) Search(firstName, lastName string) ([]*models.User, error) {
	rows, err := r.db.Queryx("SELECT * FROM users WHERE first_name LIKE ? AND second_name LIKE ? ORDER BY id",
		firstName+"%", lastName+"%")
	if err != nil {
		return nil, err
	}

	var users []*models.User
	for rows.Next() {
		var u models.User
		err = rows.StructScan(&u)
		if err != nil {
			return nil, err
		}

		users = append(users, &u)
	}

	return users, nil
}
