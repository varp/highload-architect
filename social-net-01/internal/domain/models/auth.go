package models

import "time"

type ApiKey struct {
	Key             string    `db:"api_key"`
	UserId          string    `db:"user_id"`
	AuthenticatedAt time.Time `db:"authenticated_at"`
}

type ApiKeyRepository interface {
	Create(userId string) error
	Get(key string) (*ApiKey, error)
}

func (a *ApiKey) Expired() bool {
	return a.AuthenticatedAt.Add(1 * time.Hour).Before(time.Now())
}
