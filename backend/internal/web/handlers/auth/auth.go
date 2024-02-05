package auth

import (
	"fmt"
	"net/http"

	"auction-system/internal/entity"
	"auction-system/pkg/models"
	"auction-system/pkg/utils"

	"github.com/BoryslavGlov/logrusx"

	"github.com/gin-gonic/gin"
)

type signUpRequestBody struct {
	FullName string `json:"fullName" binding:"required" example:"Bebra Bronyslavovich"`
	Username string `json:"username" binding:"required" example:"bebraslav"`
	Password string `json:"password" binding:"required" example:"mycoolpassword"`
}

func (h *Handler) SignUp(ctx *gin.Context) {
	var input signUpRequestBody

	if err := ctx.BindJSON(&input); err != nil {
		h.logx.Error("failed to BindJSON in signUp",
			logrusx.LogField{Key: "context", Value: err},
			logrusx.LogField{Key: "request", Value: fmt.Sprintf("%+v", ctx.Request)},
		)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, entity.ErrResponse{Message: err.Error()})
		return
	}

	err := h.repo.CreateUser(&models.User{
		FullName:     input.FullName,
		Username:     input.Username,
		PasswordHash: utils.HashPassword(input.Password),
	})
	if err != nil {
		h.logx.Error("failed to CreateUser",
			logrusx.LogField{Key: "context", Value: err},
			logrusx.LogField{Key: "request", Value: fmt.Sprintf("%+v", ctx.Request)},
		)
		ctx.AbortWithStatusJSON(http.StatusConflict, entity.ErrResponse{Message: "username already exists"})
		return
	}

	ctx.JSON(http.StatusCreated, entity.ErrResponse{Message: "success"})
}

type signInRequestBody struct {
	Username string `json:"username" binding:"required" example:"bebraslav"`
	Password string `json:"password" binding:"required" example:"mycoolpassword"`
}

func (h *Handler) SignIn(ctx *gin.Context) {
	var input signInRequestBody

	if err := ctx.BindJSON(&input); err != nil {
		h.logx.Error("failed to BindJSON in signIn",
			logrusx.LogField{Key: "context", Value: err},
			logrusx.LogField{Key: "request", Value: fmt.Sprintf("%+v", ctx.Request)},
		)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, entity.ErrResponse{Message: err.Error()})
		return
	}

	user, err := h.repo.GetUserByUsername(input.Username)
	if err != nil || user == nil {
		h.logx.Error("failed to GetUserByUsername",
			logrusx.LogField{Key: "context", Value: "E-Mail doesn't exist"},
			logrusx.LogField{Key: "email", Value: input.Username},
		)
		ctx.AbortWithStatusJSON(http.StatusConflict, entity.ErrResponse{Message: "username doesn't exists"})
		return
	}

	match := utils.CheckPasswordHash(input.Password, user.PasswordHash)
	if !match {
		ctx.AbortWithStatusJSON(http.StatusNotFound, entity.ErrResponse{Message: "username or password is incorrect"})
		return
	}

	accessToken, err := h.jwt.GenerateToken(*user)
	if err != nil {
		h.logx.Error("failed to GenerateToken",
			logrusx.LogField{Key: "userId", Value: user.Id.String()},
			logrusx.LogField{Key: "request", Value: fmt.Sprintf("%+v", ctx.Request)},
		)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, entity.ErrResponse{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"token": accessToken,
	})
}
