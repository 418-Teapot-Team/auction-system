package routes

import (
	"auction-system/internal/web/handlers/bids"
	"auction-system/pkg/middlewares"

	"github.com/gin-gonic/gin"
)

func BidsRouters(router *gin.RouterGroup, h *bids.Handler, middleware middlewares.Middlewares) {
	_bids := router.Group("/bids", middleware.AuthRequired)
	{
		_bids.GET("/:id", h.GetAllBids)
		// _auction.GET("/draft", h.AuctionsDraft)
	}

	// unauthorized := router.Group("/auction")
	// {
		// unauthorized.GET("/all", h.GetAllAuctions)
		// unauthorized.GET("/:id", h.GetAuction)
	// }
}