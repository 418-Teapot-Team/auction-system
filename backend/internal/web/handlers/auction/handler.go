package auction

import (
	"auction-system/pkg/repository/auction"
	"github.com/BoryslavGlov/logrusx"
)

type Handler struct {
	logx logrusx.Logging
	repo auction.Repository
}

func NewHandler(logx logrusx.Logging, repo auction.Repository) *Handler {
	return &Handler{logx: logx, repo: repo}
}
