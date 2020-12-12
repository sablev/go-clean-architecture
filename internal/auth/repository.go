package auth

import (
	"context"
	"github.com/sablev/go-clean-architecture-std/internal/models"
)

type Repository interface {
	Create(ctx context.Context, user *models.User) error
	Get(ctx context.Context, username, password string) (*models.User, error)
}
