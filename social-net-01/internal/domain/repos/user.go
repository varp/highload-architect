package repos

import (
	"fmt"
	"go.vardan.dev/highload-architect/social-net-01/internal/domain/models"
	"log"

	"github.com/jmoiron/sqlx"
	"go.vardan.dev/highload-architect/social-net-01/internal/domain/repos/mysql"
)

type userRepository struct {
	db *sqlx.DB
}

func NewUserRepository() models.UserRepository {
	db := mysql.Connect()

	return &userRepository{
		db: db,
	}
}

func (r *userRepository) Create(user *models.User) error {
	sql := "INSERT INTO users (id, age, biography, city, first_name, second_name, password_hash) VALUES (?, ?, ?," +
		" " +
		"?, ?, ?, ?)"
	_, err := r.db.Exec(sql, user.Id, user.Age, user.Biography, user.City, user.FirstName, user.SecondName,
		user.PasswordHash)
	if err != nil {
		return err
	}

	return nil
}

func (r *userRepository) Get(id string) (*models.User, error) {
	row, err := r.db.Queryx("SELECT * FROM users WHERE id = ?", id)
	if err != nil {
		log.Fatalf("User with id %s not found", id)
		return nil, err
	}

	if !row.Next() {
		err = fmt.Errorf("user with specified ID %s not found", id)
		return nil, err
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
	sql := "UPDATE users SET age = ?, biography = ?, city = ?, firstName = ?, secondName = ?, passwordHash = ? WHERE id = ?"
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