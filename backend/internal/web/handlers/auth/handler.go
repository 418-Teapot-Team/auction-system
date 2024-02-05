package auth

import (
	"auction-system/pkg/repository"
	"auction-system/pkg/utils"
	"github.com/BoryslavGlov/logrusx"
)

type Handler struct {
	logx logrusx.Logging
	repo repository.AuthRepository
	jwt  *utils.JwtManager
}

func NewHandler(logx logrusx.Logging, repo repository.AuthRepository, jwt *utils.JwtManager) *Handler {
	return &Handler{logx: logx, repo: repo, jwt: jwt}
}
