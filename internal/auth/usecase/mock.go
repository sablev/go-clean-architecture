package usecase

import (
	"context"
	"github.com/stretchr/testify/mock"
	"github.com/sablev/go-clean-architecture-std/internal/models"
)

type MockUseCase struct {
	mock.Mock
}

func (m *MockUseCase) SignUp(ctx context.Context, username, password string) error {
	args := m.Called(username, password)

	return args.Error(0)
}

func (m *MockUseCase) SignIn(ctx context.Context, username, password string) (string, error) {
	args := m.Called(username, password)

	return args.Get(0).(string), args.Error(1)
}

func (m *MockUseCase) ParseToken(ctx context.Context, accessToken string) (*models.User, error) {
	args := m.Called(accessToken)

	return args.Get(0).(*models.User), args.Error(1)
}