package auth

import (
	"context"
	"github.com/sablev/go-clean-architecture/internal/models"
)

const CtxUserKey = "user"

type UseCase interface {
	SignUp(ctx context.Context, username, password string) error
	SignIn(ctx context.Context, username, password string) (string, error)
	ParseToken(ctx context.Context, accessToken string) (*models.User, error)
}
