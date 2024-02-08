package users

import (
	"auction-system/internal/entity"
	"auction-system/pkg/middlewares"
	"github.com/BoryslavGlov/logrusx"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) WhoAmI(ctx *gin.Context) {
	userId := middlewares.GetUserId(ctx)

	user, err := h.repo.GetUserById(userId)
	if err != nil || user == nil {
		h.logx.Error("failed to GetUserByUsername",
			logrusx.LogField{Key: "context", Value: "user id doesn't exist"},
			logrusx.LogField{Key: "id", Value: userId},
		)
		ctx.AbortWithStatusJSON(http.StatusConflict, entity.ErrResponse{Message: "id doesn't exists"})
		return
	}

	ctx.JSON(http.StatusOK, user)
}
