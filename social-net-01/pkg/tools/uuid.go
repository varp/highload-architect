package tools

import (
	"github.com/google/uuid"
	"go.vardan.dev/highload-architect/social-net-01/internal/domain/models"
)

type UUIDGenerator struct{}

func NewUUIDGenerator() models.UUIDGenerator {
	return &UUIDGenerator{}
}

func (g UUIDGenerator) Generate() string {
	return uuid.NewString()
}
