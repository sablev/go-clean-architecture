package http

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/sablev/go-clean-architecture-std/internal/auth"
	"github.com/sablev/go-clean-architecture-std/internal/auth/usecase"
	"github.com/sablev/go-clean-architecture-std/internal/models"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAuthMiddleware(t *testing.T) {
	r := gin.Default()
	uc := new(usecase.MockUseCase)

	r.POST("/api/endpoint", NewMiddleware(uc), func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	w := httptest.NewRecorder()

	// No Auth Header request
	req, _ := http.NewRequest("POST", "/api/endpoint", nil)
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusUnauthorized, w.Code)

	// Empty Auth Header request
	w = httptest.NewRecorder()

	req.Header.Set("Authorization", "")
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusUnauthorized, w.Code)

	// Bearer Auth Header with no token request
	w = httptest.NewRecorder()
	uc.On("ParseToken", "").Return(&models.User{}, auth.ErrInvalidAccessToken)

	req.Header.Set("Authorization", "Bearer ")
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusUnauthorized, w.Code)

	// Valid Auth Header
	w = httptest.NewRecorder()
	uc.On("ParseToken", "token").Return(&models.User{}, nil)

	req.Header.Set("Authorization", "Bearer token")
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}
