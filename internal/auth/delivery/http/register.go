package http

import (
	"github.com/gin-gonic/gin"
	"github.com/sablev/go-clean-architecture-std/internal/auth"
)

func RegisterHTTPEndpoints(router *gin.Engine, uc auth.UseCase) {
	h := NewHandler(uc)

	authEndpoints := router.Group("/auth")
	{
		authEndpoints.POST("/sign-up", h.SignUp)
		authEndpoints.POST("/sign-in", h.SignIn)
	}
}
