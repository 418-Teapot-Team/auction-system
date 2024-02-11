package bids

import (
	"auction-system/pkg/middlewares"
	"auction-system/pkg/models"
	"auction-system/pkg/repository/bids"
	"encoding/json"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

type Message struct {
	Type    string          `json:"type"`
	Payload json.RawMessage `json:"payload"`
}

type NewBidPayload struct {
	AuctionId string     `json:"auction_id"`
	NewValue  int64 `json:"new_value"`
}

func (h *Handler) GetAllBids(ctx *gin.Context) {
	conn, err := h.upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		log.Println("Error upgrading to WebSocket:", err)
		return
	}

	defer conn.Close()

	// Register connection
	h.connections[conn] = true
	defer delete(h.connections, conn)

	h.logx.Info("New WebSocket connection")
	for {
		var msg Message
		if err := conn.ReadJSON(&msg); err != nil {
			log.Println("Error reading message:", err)
			break
		}
		switch msg.Type {
			case "bid": handleBid(h.repo, msg.Payload, ctx)
			default: h.logx.Info(fmt.Sprintf("Unknown message type: %s", msg.Type))
		}
		h.logx.Info(fmt.Sprintf("Received message: %s", msg.Type))
	}
}

func handleBid(repo bids.Repository, payload json.RawMessage, ctx *gin.Context) {

	var newBidPayload NewBidPayload
	if err := json.Unmarshal(payload, &newBidPayload); err != nil {
		log.Println("Error unmarshalling payload:", err)
		return
	}

	userId := middlewares.GetUserId(ctx)

	bid := &models.Bid{
		BidderId: userId,
		AuctionId: newBidPayload.AuctionId,
		NewValue:  newBidPayload.NewValue,
	}

	if err := repo.CreateBid(bid); err != nil {
		log.Println("Error creating bid:", err)
		return
	}
}