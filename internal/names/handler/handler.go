package handler

import (
	"TestTaskEffectiveMobile/internal/names/model"
	"TestTaskEffectiveMobile/pkg/logger"
	"context"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Service interface {
	CreateUser(ctx context.Context, p model.PersonRequest) error
}

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

func (h Handler) CreateUser(contx context.Context) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var p model.PersonRequest
		if err := ctx.ShouldBindJSON(&p); err != nil {
			logger.GetLoggerFromCtx(contx).Info(contx, "Failed bind json", zap.Error(err))
			ctx.JSON(400, "Failed bind json")
			return
		}

		err := h.service.CreateUser(contx, p)
		if err != nil {
			logger.GetLoggerFromCtx(contx).Info(contx, "Failed Create new person", zap.Error(err))
			ctx.JSON(500, "Failed Create new person")
			return
		}

		ctx.JSON(200, "Succesfully created new person")
	}
}
