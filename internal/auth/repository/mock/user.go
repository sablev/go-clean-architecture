package mock

import (
	"context"
	"github.com/sablev/go-clean-architecture-std/internal/models"
	"github.com/stretchr/testify/mock"
)

type UserStorageMock struct {
	mock.Mock
}

func (s *UserStorageMock) Create(ctx context.Context, user *models.User) error {
	args := s.Called(user)

	return args.Error(0)
}

func (s *UserStorageMock) Get(ctx context.Context, username, password string) (*models.User, error) {
	args := s.Called(username, password)

	return args.Get(0).(*models.User), args.Error(1)
}
