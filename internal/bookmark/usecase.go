package bookmark

import (
	"context"

	"github.com/sablev/go-clean-architecture-std/internal/models"
)

type UseCase interface {
	Create(ctx context.Context, user *models.User, url, title string) error
	Get(ctx context.Context, user *models.User) ([]*models.Bookmark, error)
	Delete(ctx context.Context, user *models.User, id string) error
}
