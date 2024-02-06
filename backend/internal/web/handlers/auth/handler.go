package auth

import (
	"auction-system/pkg/repository/auth"
	"auction-system/pkg/utils"
	"github.com/BoryslavGlov/logrusx"
)

type Handler struct {
	logx logrusx.Logging
	repo auth.Repository
	jwt  *utils.JwtManager
}

func NewHandler(logx logrusx.Logging, repo auth.Repository, jwt *utils.JwtManager) *Handler {
	return &Handler{logx: logx, repo: repo, jwt: jwt}
}
