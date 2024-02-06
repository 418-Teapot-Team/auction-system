package routes

import (
	"auction-system/internal/web/handlers/auction"
	"auction-system/pkg/middlewares"
	"github.com/gin-gonic/gin"
)

func AuctionRouters(router *gin.RouterGroup, h *auction.Handler, middleware middlewares.Middlewares) {
	_auction := router.Group("/auction", middleware.AuthRequired)
	{
		_auction.POST("/create", h.CreateAuction)
	}
}
