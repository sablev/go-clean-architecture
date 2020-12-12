package mock

import (
	"context"
	"github.com/sablev/go-clean-architecture-std/internal/models"
	"github.com/stretchr/testify/mock"
)

type Storage struct {
	mock.Mock
}

func (s *Storage) Create(ctx context.Context, user *models.User) error {
	args := s.Called(user)

	return args.Error(0)
}

func (s *Storage) Get(ctx context.Context, username, password string) (*models.User, error) {
	args := s.Called(username, password)

	return args.Get(0).(*models.User), args.Error(1)
}
