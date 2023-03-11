package user

import (
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"go.vardan.dev/highload-architect/social-net-01/internal/domain/models/user"
)

type repository struct {
	db *sqlx.DB
}

func New() user.Repository {
	// Capture connection properties.
	cfg := mysql.Config{
		User:   os.Getenv("SOCIAL_NET_01_DB_USER"), // root
		Passwd: os.Getenv("SOCIAL_NET_01_DB_PASS"), // pass
		Net:    os.Getenv("SOCIAL_NET_01_DB_NET"),  // tcp
		Addr:   os.Getenv("SOCIAL_NET_01_DB_ADDR"), // 127.0.0.1:3306
		DBName: os.Getenv("SOCIAL_NET_01_DB_NAME"), // social-net-01
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
	return nil
}

func (r *repository) Get(id string) (*user.User, error) {
	row, err := r.db.Queryx("SELECT * FROM user WHERE uid = ?", id)
	if err != nil {
		log.Fatalf("User with id %s not found", id)
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
	return nil
}

func (r *repository) Delete(user *user.User) error {
	return nil
}
