package localcache

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/sablev/go-clean-architecture-std/internal/bookmark"
	"github.com/sablev/go-clean-architecture-std/internal/models"
	"testing"
)

func TestGetBookmarks(t *testing.T) {
	id := "id"
	user := &models.User{ID: id}

	s := NewBookmarkLocalStorage()

	for i := 0; i < 10; i++ {
		bm := &models.Bookmark{
			ID:     fmt.Sprintf("id%d", i),
			UserID: user.ID,
		}

		err := s.Create(context.Background(), user, bm)
		assert.NoError(t, err)
	}

	returnedBookmarks, err := s.Get(context.Background(), user)
	assert.NoError(t, err)

	assert.Equal(t, 10, len(returnedBookmarks))
}

func TestDeleteBookmark(t *testing.T) {
	id1 := "id1"
	id2 := "id2"

	user1 := &models.User{ID: id1}
	user2 := &models.User{ID: id2}

	bmID := "bmID"
	bm := &models.Bookmark{ID: bmID, UserID: user1.ID}

	s := NewBookmarkLocalStorage()

	err := s.Create(context.Background(), user1, bm)
	assert.NoError(t, err)

	err = s.Delete(context.Background(), user1, bmID)
	assert.NoError(t, err)

	err = s.Create(context.Background(), user1, bm)
	assert.NoError(t, err)

	err = s.Delete(context.Background(), user2, bmID)
	assert.Error(t, err)
	assert.Equal(t, err, bookmark.ErrBookmarkNotFound)
}
