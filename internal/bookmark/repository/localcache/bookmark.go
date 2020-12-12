package localcache

import (
	"context"
	"sync"

	"github.com/sablev/go-clean-architecture-std/internal/bookmark"
	"github.com/sablev/go-clean-architecture-std/internal/models"
)

type Storage struct {
	bookmarks map[string]*models.Bookmark
	mutex     *sync.Mutex
}

func New() *Storage {
	return &Storage{
		bookmarks: make(map[string]*models.Bookmark),
		mutex:     new(sync.Mutex),
	}
}

func (s *Storage) Create(ctx context.Context, user *models.User, bm *models.Bookmark) error {
	bm.UserID = user.ID

	s.mutex.Lock()
	s.bookmarks[bm.ID] = bm
	s.mutex.Unlock()

	return nil
}

func (s *Storage) Get(ctx context.Context, user *models.User) ([]*models.Bookmark, error) {
	bookmarks := make([]*models.Bookmark, 0)

	s.mutex.Lock()
	for _, bm := range s.bookmarks {
		if bm.UserID == user.ID {
			bookmarks = append(bookmarks, bm)
		}
	}
	s.mutex.Unlock()

	return bookmarks, nil
}

func (s *Storage) Delete(ctx context.Context, user *models.User, id string) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	bm, ex := s.bookmarks[id]
	if ex && bm.UserID == user.ID {
		delete(s.bookmarks, id)
		return nil
	}

	return bookmark.ErrBookmarkNotFound
}
