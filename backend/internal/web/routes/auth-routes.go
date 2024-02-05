package routes

import (
	"auction-system/internal/web/handlers/auth"
	"github.com/gin-gonic/gin"
)

func AuthRouters(router *gin.RouterGroup, h *auth.Handler) {
	_auth := router.Group("/auth")
	{
		_auth.GET("/signup", h.SignUp)
	}
}
