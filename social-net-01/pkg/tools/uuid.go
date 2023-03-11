package tools

import "github.com/google/uuid"

type UUIDGenerator struct{}

func (g UUIDGenerator) Generate() string {
	return uuid.NewString()
}
