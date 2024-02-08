package users

import (
	"auction-system/pkg/repository/users"
	"github.com/BoryslavGlov/logrusx"
)

type Handler struct {
	logx logrusx.Logging
	repo users.Repository
}

func NewHandler(logx logrusx.Logging, repo users.Repository) *Handler {
	return &Handler{logx: logx, repo: repo}
}
