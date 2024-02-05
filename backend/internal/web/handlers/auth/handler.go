package auth

import (
	"auction-system/pkg/repository"
	"github.com/BoryslavGlov/logrusx"
)

type Handler struct {
	logx logrusx.Logging
	repo repository.AuthRepository
}

func NewHandler(logx logrusx.Logging, repo repository.AuthRepository) *Handler {
	return &Handler{logx: logx, repo: repo}
}
