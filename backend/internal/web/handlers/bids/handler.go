package bids

import (
	"auction-system/pkg/repository/auction"
	"auction-system/pkg/repository/bids"

	"github.com/BoryslavGlov/logrusx"
	"github.com/gorilla/websocket"
)


type Handler struct {
	logx logrusx.Logging
	repo bids.Repository
	auctionRepo auction.Repository
	upgrader websocket.Upgrader
	connections map[*websocket.Conn]bool
}

func NewHandler(logx logrusx.Logging, repo bids.Repository, auctionRepo auction.Repository, upgrader websocket.Upgrader, connections map[*websocket.Conn]bool) *Handler {
	return &Handler{logx: logx, repo: repo, upgrader: upgrader, connections: connections}
}