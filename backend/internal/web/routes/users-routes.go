package routes

import (
	"auction-system/internal/web/handlers/users"
	"auction-system/pkg/middlewares"
	"github.com/gin-gonic/gin"
)

func UsersRouters(router *gin.RouterGroup, h *users.Handler, middleware middlewares.Middlewares) {
	_users := router.Group("/users", middleware.AuthRequired)
	{
		_users.GET("/me", h.WhoAmI)
	}
}
