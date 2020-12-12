package mock

import (
	"context"
	"github.com/stretchr/testify/mock"
	"github.com/sablev/go-clean-architecture-std/internal/models"
)

type Storage struct {
	mock.Mock
}

func (s *Storage) Create(ctx context.Context, user *models.User, bm *models.Bookmark) error {
	args := s.Called(user, bm)

	return args.Error(0)
}

func (s *Storage) Get(ctx context.Context, user *models.User) ([]*models.Bookmark, error) {
	args := s.Called(user)

	return args.Get(0).([]*models.Bookmark), args.Error(1)
}

func (s *Storage) Delete(ctx context.Context, user *models.User, id string) error {
	args := s.Called(user, id)

	return args.Error(0)
}
