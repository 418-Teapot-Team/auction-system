package auction

import (
	"fmt"
	"net/http"

	"auction-system/internal/entity"
	"auction-system/pkg/middlewares"
	"auction-system/pkg/models"

	"github.com/BoryslavGlov/logrusx"
	"github.com/gin-gonic/gin"
)

type createAuctionRequestBody struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	StartBit    int64  `json:"startBit" binding:"required"`
}

func (h *Handler) CreateAuction(ctx *gin.Context) {
	var input createAuctionRequestBody

	userId := middlewares.GetUserId(ctx)

	if err := ctx.BindJSON(&input); err != nil {
		h.logx.Error("failed to BindJSON CreateAuction",
			logrusx.LogField{Key: "context", Value: err},
			logrusx.LogField{Key: "request", Value: fmt.Sprintf("%+v", ctx.Request)},
		)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, entity.ErrResponse{Message: err.Error()})
		return
	}

	auction := &models.Auction{
		CreatorId:   userId,
		Title:       input.Title,
		Description: input.Description,
		StartBit:    input.StartBit,
		CurrentBit:  input.StartBit,
	}

	err := h.repo.CreateAuction(auction)
	if err != nil {
		h.logx.Error("failed to createAuction",
			logrusx.LogField{Key: "context", Value: err},
			logrusx.LogField{Key: "request", Value: fmt.Sprintf("%+v", ctx.Request)},
		)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, entity.ErrResponse{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"id": auction.Id.String(),
	})
}
