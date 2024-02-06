package auction

import (
	"fmt"
	"net/http"
	"time"

	"auction-system/internal/entity"
	"auction-system/pkg/middlewares"
	"auction-system/pkg/models"

	"github.com/google/uuid"

	"github.com/BoryslavGlov/logrusx"
	"github.com/gin-gonic/gin"
)

type createAuctionRequestBody struct {
	Title       string   `json:"title" binding:"required"`
	Description string   `json:"description" binding:"required"`
	StartBit    int64    `json:"startBit" binding:"required"`
	Images      []string `json:"images" binding:"required"`
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

	var images []models.Images

	for _, image := range input.Images {
		images = append(images, models.Images{
			DownloadUrl: image,
		})
	}

	auction := &models.Auction{
		CreatorId:   userId,
		Title:       input.Title,
		Description: input.Description,
		StartBit:    input.StartBit,
		CurrentBit:  input.StartBit,
		Images:      images,
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

func (h *Handler) GetAllAuctions(ctx *gin.Context) {
	auctions, err := h.repo.GetAllAuctions()
	if err != nil {
		h.logx.Error("failed to getAllAuctions",
			logrusx.LogField{Key: "context", Value: err},
			logrusx.LogField{Key: "request", Value: fmt.Sprintf("%+v", ctx.Request)},
		)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, entity.ErrResponse{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": auctions})
}

func (h *Handler) GetAuction(ctx *gin.Context) {
	auctionId := ctx.Param("id")

	auction, err := h.repo.GetAuctionById(auctionId)
	if err != nil {
		h.logx.Error("failed to getAuctionById",
			logrusx.LogField{Key: "context", Value: err},
			logrusx.LogField{Key: "request", Value: fmt.Sprintf("%+v", ctx.Request)},
		)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, entity.ErrResponse{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, auction)
}

type auctionDraftResponse struct {
	Id          *uuid.UUID      `gorm:"column:id;->" json:"id"`
	Title       string          `gorm:"column:title" json:"title"`
	Description string          `gorm:"column:description" json:"description"`
	StartBit    int64           `gorm:"column:startbit" json:"startBit"`
	CurrentBit  int64           `gorm:"column:currentbit" json:"currentBit"`
	CreatedAt   time.Time       `gorm:"column:createdat;->" json:"createdAt"`
	UpdatedAt   time.Time       `gorm:"column:updatedat" json:"updatedAt"`
	Images      []models.Images `gorm:"foreignKey:AuctionId;" json:"images"`
}

func (h *Handler) AuctionsDraft(ctx *gin.Context) {
	userId := middlewares.GetUserId(ctx)

	auctions, err := h.repo.GetAuctionsByUserId(userId)
	if err != nil {
		h.logx.Error("failed to GetAuctionByUserId",
			logrusx.LogField{Key: "context", Value: err},
			logrusx.LogField{Key: "request", Value: fmt.Sprintf("%+v", ctx.Request)},
		)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, entity.ErrResponse{Message: err.Error()})
		return
	}

	var draft []auctionDraftResponse

	for _, auction := range auctions {
		draft = append(draft, auctionDraftResponse{
			Id:          auction.Id,
			Title:       auction.Title,
			Description: auction.Description,
			StartBit:    auction.StartBit,
			CurrentBit:  auction.CurrentBit,
			CreatedAt:   auction.CreatedAt,
			UpdatedAt:   auction.UpdatedAt,
			Images:      auction.Images,
		})
	}

	ctx.JSON(http.StatusOK, gin.H{"data": draft})
}
