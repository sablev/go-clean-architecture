package usecase

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/sablev/go-clean-architecture-std/internal/auth/repository/mock"
	"github.com/sablev/go-clean-architecture-std/internal/models"
	"testing"
)

func TestAuthFlow(t *testing.T) {
	repo := new(mock.Storage)

	uc := New(repo, "salt", []byte("secret"), 86400)

	var (
		username = "user"
		password = "pass"

		ctx = context.Background()

		user = &models.User{
			Username: username,
			Password: "11f5639f22525155cb0b43573ee4212838c78d87", // sha1 of pass+salt
		}
	)

	// Sign Up
	repo.On("Create", user).Return(nil)
	err := uc.SignUp(ctx, username, password)
	assert.NoError(t, err)

	// Sign In (Get Auth Token)
	repo.On("Get", user.Username, user.Password).Return(user, nil)
	token, err := uc.SignIn(ctx, username, password)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	// Verify token
	parsedUser, err := uc.ParseToken(ctx, token)
	assert.NoError(t, err)
	assert.Equal(t, user, parsedUser)
}
