package user

import (
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"go.vardan.dev/highload-architect/social-net-01/internal/domain/models/user"
	"go.vardan.dev/highload-architect/social-net-01/pkg/tools"
)

type repository struct {
	db *sqlx.DB
}

func New() user.Repository {
	env := tools.NewEnv("")

	// Capture connection properties.
	cfg := mysql.Config{
		Addr: env.GetWithDefault("DB_ADDR",
			fmt.Sprintf("%s:%s",
				env.GetWithDefault("DB_HOST", "localhost"),
				env.GetWithDefault("DB_PORT", "3306"),
			),
		),
		User:   env.GetWithDefault("DB_USER", "root"),
		Passwd: env.GetWithDefault("DB_PASS", "pass"),
		Net:    env.GetWithDefault("DB_NET", "tcp"),
		DBName: env.GetWithDefault("DB_NAME", "social-net"),
	}
	// Get a database handle.
	var err error
	var db *sqlx.DB

	db, err = sqlx.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")

	return &repository{
		db: db,
	}
}

func (r *repository) Create(user *user.User) error {
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

func (r *repository) Get(id string) (*user.User, error) {
	row, err := r.db.Queryx("SELECT * FROM users WHERE uid = ?", id)
	if err != nil {
		log.Fatalf("User with id %s not found", id)
		return nil, err
	}

	if !row.Next() {
		err = fmt.Errorf("user with specified ID %s not found", id)
		return nil, err
	}

	var u user.User
	err = row.StructScan(&u)
	if err != nil {
		log.Fatalf("Failed to map DB response to domain User model")
		return nil, err
	}

	return &u, nil
}

func (r *repository) Update(user *user.User) error {
	sql := "UPDATE users SET age = ?, biography = ?, city = ?, firstName = ?, secondName = ?, passwordHash = ? WHERE id = ?"
	_, err := r.db.Exec(sql, user.Age, user.Biography, user.City, user.FirstName, user.SecondName,
		user.PasswordHash, user.Id)
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) Delete(user *user.User) error {
	_, err := r.db.Exec("DELETE FROM users WHERE id = ?", user.Id)
	if err != nil {
		return err
	}

	return nil
}
