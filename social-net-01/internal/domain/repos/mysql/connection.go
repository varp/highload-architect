package mysql

import (
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"go.vardan.dev/highload-architect/social-net-01/pkg/tools"
)

func Connect() *sqlx.DB {
	env := tools.NewEnv()

	// Capture connection properties.
	cfg := mysql.Config{
		Addr: env.GetWithDefault("DB_ADDR",
			fmt.Sprintf("%s:%s",
				env.GetWithDefault("DB_HOST", "localhost"),
				env.GetWithDefault("DB_PORT", "3306"),
			),
		),
		User:      env.GetWithDefault("DB_USER", "root"),
		Passwd:    env.GetWithDefault("DB_PASS", "pass"),
		Net:       env.GetWithDefault("DB_NET", "tcp"),
		DBName:    env.GetWithDefault("DB_NAME", "social-net"),
		ParseTime: true,
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

	log.Println("Connected!")

	return db
}
