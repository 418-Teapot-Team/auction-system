package middlewares

import (
	"net/http"
	"strings"

	"auction-system/internal/entity"
	"auction-system/pkg/utils"

	"github.com/gin-gonic/gin"
)

const (
	userCtx       = "userId"
	authorization = "Authorization"
)

type Middlewares interface {
	AuthRequired(ctx *gin.Context)
	CORSMiddleware() gin.HandlerFunc
}

type middleware struct {
	svc *utils.JwtManager
}

func NewMiddlewareService(svc *utils.JwtManager) Middlewares {
	return &middleware{svc: svc}
}

func (c *middleware) AuthRequired(ctx *gin.Context) {
	authHeader := ctx.Request.Header.Get(authorization)

	if authHeader == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, entity.ErrResponse{Message: "missing authorization header in request"})
		return
	}

	token := strings.Split(authHeader, "Bearer ")

	if len(token) < 2 {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, entity.ErrResponse{Message: "invalid header value"})
		return
	}

	if token[1] == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, entity.ErrResponse{Message: "invalid header value"})
		return
	}

	claims, err := c.svc.ValidateToken(token[1])
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, entity.ErrResponse{Message: err.Error()})
		return
	}

	ctx.Set(userCtx, claims.Subject)

	ctx.Next()
}

func GetUserId(ctx *gin.Context) string {
	userId, ok := ctx.Get(userCtx)
	if !ok {
		return ""
	}
	return userId.(string)
}
