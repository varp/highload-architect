package repos

import (
	"time"

	"github.com/jmoiron/sqlx"
	"go.vardan.dev/highload-architect/social-net-01/internal/domain/models"
	"go.vardan.dev/highload-architect/social-net-01/internal/domain/repos/mysql"
)

type apiKeyRepository struct {
	db            *sqlx.DB
	uuidGenerator models.UUIDGenerator
}

func NewApiKeyRepository(uuidGenerator models.UUIDGenerator) models.ApiKeyRepository {
	db := mysql.Connect() // TODO: pass db connection as a parameter

	return &apiKeyRepository{
		db:            db,
		uuidGenerator: uuidGenerator,
	}
}

func (a *apiKeyRepository) Create(userId string) (*models.ApiKey, error) {
	key := a.uuidGenerator.Generate()

	sql := "INSERT INTO `auth` (`api_key`, `user_id`, `authenticated_at`)  VALUES (?, ?, ?)"
	// TODO: research how to use timezones in MySQL and Go
	// see: https://dev.mysql.com/doc/refman/8.0/en/time-zone-support.html
	// see: https://golang.org/pkg/time/#Time.Format
	authenticatedAt := time.Now().UTC()

	_, err := a.db.Exec(sql, key, userId, authenticatedAt.Format(time.DateTime))
	if err != nil {
		return nil, err
	}

	return &models.ApiKey{Key: key, UserId: userId, AuthenticatedAt: authenticatedAt}, nil
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

func (a *apiKeyRepository) GetByUserId(userId string) (*models.ApiKey, error) {
	row, err := a.db.Queryx("SELECT * FROM `auth` WHERE user_id = ?", userId)
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

func (a *apiKeyRepository) Delete(key string) error {
	sql := "DELETE FROM `auth` WHERE api_key = ?"
	_, err := a.db.Exec(sql, key)
	if err != nil {
		return err
	}

	return nil
}
