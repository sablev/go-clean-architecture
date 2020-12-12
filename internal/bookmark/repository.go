package bookmark

import (
	"context"

	"github.com/sablev/go-clean-architecture-std/internal/models"
)

type Repository interface {
	Create(ctx context.Context, user *models.User, bm *models.Bookmark) error
	Get(ctx context.Context, user *models.User) ([]*models.Bookmark, error)
	Delete(ctx context.Context, user *models.User, id string) error
}
