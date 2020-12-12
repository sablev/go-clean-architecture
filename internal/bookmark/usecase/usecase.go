package usecase

import (
	"context"

	"github.com/sablev/go-clean-architecture-std/internal/bookmark"
	"github.com/sablev/go-clean-architecture-std/internal/models"
)

type UseCase struct {
	bookmarkRepo bookmark.Repository
}

func New(bookmarkRepo bookmark.Repository) *UseCase {
	return &UseCase{
		bookmarkRepo: bookmarkRepo,
	}
}

func (b UseCase) Create(ctx context.Context, user *models.User, url, title string) error {
	bm := &models.Bookmark{
		URL:   url,
		Title: title,
	}

	return b.bookmarkRepo.Create(ctx, user, bm)
}

func (b UseCase) Get(ctx context.Context, user *models.User) ([]*models.Bookmark, error) {
	return b.bookmarkRepo.Get(ctx, user)
}

func (b UseCase) Delete(ctx context.Context, user *models.User, id string) error {
	return b.bookmarkRepo.Delete(ctx, user, id)
}
