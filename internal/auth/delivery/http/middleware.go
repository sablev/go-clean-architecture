package http

import (
	"github.com/gin-gonic/gin"
	"github.com/sablev/go-clean-architecture-std/internal/auth"
	"net/http"
	"strings"
)

type Middleware struct {
	usecase auth.UseCase
}

func NewMiddleware(usecase auth.UseCase) gin.HandlerFunc {
	return (&Middleware{
		usecase: usecase,
	}).Handle
}

func (m *Middleware) Handle(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	headerParts := strings.Split(authHeader, " ")
	if len(headerParts) != 2 {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	if headerParts[0] != "Bearer" {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	user, err := m.usecase.ParseToken(c.Request.Context(), headerParts[1])
	if err != nil {
		status := http.StatusInternalServerError
		if err == auth.ErrInvalidAccessToken {
			status = http.StatusUnauthorized
		}

		c.AbortWithStatus(status)
		return
	}

	c.Set(auth.CtxUserKey, user)
}
