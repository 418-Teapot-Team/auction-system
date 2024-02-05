package auth

import (
	"auction-system/pkg/repository"
	"github.com/BoryslavGlov/logrusx"
)

type Handler struct {
	logx logrusx.Logging
	repo repository.Repository
}

func NewHandler(logx logrusx.Logging, repo repository.Repository) *Handler {
	return &Handler{logx: logx, repo: repo}
}
