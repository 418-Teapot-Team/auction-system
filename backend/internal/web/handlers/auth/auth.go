package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) SignUp(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "hello world"})
}
