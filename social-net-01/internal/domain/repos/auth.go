package repos

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"go.vardan.dev/highload-architect/social-net-01/internal/domain/models"
	"go.vardan.dev/highload-architect/social-net-01/internal/domain/repos/mysql"
)

type apiKeyRepository struct {
	db *sqlx.DB
}

func NewApiKeyRepository() models.ApiKeyRepository {
	db := mysql.Connect()

	return &apiKeyRepository{
		db: db,
	}
}

func (a *apiKeyRepository) Create(userId string) error {

	key := uuid.New().String()

	sql := "INSERT INTO `auth` (`api_key`, `user_id`)  VALUES (?, ?)"
	_, err := a.db.Exec(sql, key, userId)

	if err != nil {
		return err
	}

	return nil
}

func (a *apiKeyRepository) Get(key string) (*models.ApiKey, error) {

	row, err := a.db.Queryx("SELECT * FROM `auth` WHERE api_key = ?", key)
	if err != nil {
		return nil, err
	}

	if !row.Next() {
		return nil, nil
	}

	var am models.ApiKey
	err = row.StructScan(&am)
	if err != nil {
		return nil, err
	}

	return &am, nil
}
