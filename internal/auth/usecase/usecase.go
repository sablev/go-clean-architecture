package usecase

import (
	"context"
	"crypto/sha1"
	"fmt"
	"time"

	"github.com/sablev/go-clean-architecture-std/internal/auth"
	"github.com/sablev/go-clean-architecture-std/internal/models"

	"github.com/dgrijalva/jwt-go/v4"
)

type AuthClaims struct {
	jwt.StandardClaims
	User *models.User `json:"user"`
}

type UseCase struct {
	userRepo       auth.Repository
	hashSalt       string
	signingKey     []byte
	expireDuration time.Duration
}

func New(
	userRepo auth.Repository,
	hashSalt string,
	signingKey []byte,
	tokenTTLSeconds time.Duration) *UseCase {
	return &UseCase{
		userRepo:       userRepo,
		hashSalt:       hashSalt,
		signingKey:     signingKey,
		expireDuration: time.Second * tokenTTLSeconds,
	}
}

func (a *UseCase) SignUp(ctx context.Context, username, password string) error {
	pwd := sha1.New()
	pwd.Write([]byte(password))
	pwd.Write([]byte(a.hashSalt))

	user := &models.User{
		Username: username,
		Password: fmt.Sprintf("%x", pwd.Sum(nil)),
	}

	return a.userRepo.Create(ctx, user)
}

func (a *UseCase) SignIn(ctx context.Context, username, password string) (string, error) {
	pwd := sha1.New()
	pwd.Write([]byte(password))
	pwd.Write([]byte(a.hashSalt))
	password = fmt.Sprintf("%x", pwd.Sum(nil))

	user, err := a.userRepo.Get(ctx, username, password)
	if err != nil {
		return "", auth.ErrUserNotFound
	}

	claims := AuthClaims{
		User: user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwt.At(time.Now().Add(a.expireDuration)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(a.signingKey)
}

func (a *UseCase) ParseToken(ctx context.Context, accessToken string) (*models.User, error) {
	token, err := jwt.ParseWithClaims(accessToken, &AuthClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return a.signingKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*AuthClaims); ok && token.Valid {
		return claims.User, nil
	}

	return nil, auth.ErrInvalidAccessToken
}
