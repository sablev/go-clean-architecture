package localstorage

import (
	"context"
	"sync"

	"github.com/sablev/go-clean-architecture-std/internal/auth"
	"github.com/sablev/go-clean-architecture-std/internal/models"
)

type UserLocalStorage struct {
	users map[string]*models.User
	mutex *sync.Mutex
}

func NewUserLocalStorage() *UserLocalStorage {
	return &UserLocalStorage{
		users: make(map[string]*models.User),
		mutex: new(sync.Mutex),
	}
}

func (s *UserLocalStorage) Create(ctx context.Context, user *models.User) error {
	s.mutex.Lock()
	s.users[user.ID] = user
	s.mutex.Unlock()

	return nil
}

func (s *UserLocalStorage) Get(ctx context.Context, username, password string) (*models.User, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	for _, user := range s.users {
		if user.Username == username && user.Password == password {
			return user, nil
		}
	}

	return nil, auth.ErrUserNotFound
}
