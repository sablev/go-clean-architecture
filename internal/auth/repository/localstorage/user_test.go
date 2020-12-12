package localstorage

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/sablev/go-clean-architecture-std/internal/auth"
	"github.com/sablev/go-clean-architecture-std/internal/models"
	"testing"
)

func TestGetUser(t *testing.T) {
	s := New()

	id1 := "id"

	user := &models.User{
		ID:       id1,
		Username: "user",
		Password: "password",
	}

	err := s.Create(context.Background(), user)
	assert.NoError(t, err)

	returnedUser, err := s.Get(context.Background(), "user", "password")
	assert.NoError(t, err)
	assert.Equal(t, user, returnedUser)

	returnedUser, err = s.Get(context.Background(), "user", "")
	assert.Error(t, err)
	assert.Equal(t, err, auth.ErrUserNotFound)
}
