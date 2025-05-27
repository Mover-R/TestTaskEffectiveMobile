package handler

import (
	"TestTaskEffectiveMobile/internal/names/model"
	"TestTaskEffectiveMobile/pkg/logger"
	"context"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"
)

// @title Person Service API
// @version 1.0
// @description API for managing persons with enriched data
// @BasePath /api/v1

type Service interface {
	CreateUser(ctx context.Context, p model.PersonRequest) error
	GetUser(ctx context.Context, userID int) (model.Person, error)
	UpdateUser(ctx context.Context, userID int, p model.Person) error
	DeleteUser(ctx context.Context, userID int) error
	FindWithFilter(ctx context.Context, filter model.Filter) ([]model.Person, error)
}

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

// CreateUser godoc
// @Summary Create a new person
// @Description Create person with enriched data (age, gender, nationality)
// @Tags Persons
// @Accept json
// @Produce json
// @Param input body model.PersonRequest true "Person data"
// @Success 200 {object} string
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /api/users/create [post]
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

// GetUser godoc
// @Summary Get person by ID
// @Description Get person details by ID
// @Tags Persons
// @Accept json
// @Produce json
// @Param user_id path int true "Person ID"
// @Success 200 {object} model.Person
// @Failure 204 {object} string
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /api/users/get/{user_id} [get]
func (h Handler) GetUser(contx context.Context) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//var userID int
		userID, err := strconv.Atoi(ctx.Param("user_id"))
		if err != nil {
			logger.GetLoggerFromCtx(contx).Info(contx, "Failed param user_id", zap.Error(err))
			ctx.JSON(400, "Bad request: Failed param user_id")
			return
		}

		p, err := h.service.GetUser(contx, userID)
		if err == pgx.ErrNoRows {
			ctx.JSON(204, "No such user")
			return
		}
		if err != nil {
			logger.GetLoggerFromCtx(contx).Info(contx, "Failed get user", zap.Error(err))
			ctx.JSON(500, "Failed get user")
			return
		}

		ctx.JSON(200, p)
	}
}

// UpdateUser godoc
// @Summary Update person
// @Description Update person information
// @Tags Persons
// @Accept json
// @Produce json
// @Param user_id path int true "Person ID"
// @Param input body model.Person true "Updated person data"
// @Success 200 {object} string
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /api/users/update/{user_id} [put]
func (h Handler) UpdateUser(contx context.Context) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userID, err := strconv.Atoi(ctx.Param("user_id"))
		if err != nil {
			logger.GetLoggerFromCtx(contx).Info(contx, "Failed param user_id", zap.Error(err))
			ctx.JSON(400, "Bad request: Failed param user_id")
			return
		}

		var p model.Person
		if err := ctx.ShouldBindJSON(&p); err != nil {
			logger.GetLoggerFromCtx(contx).Info(contx, "Failed bind json", zap.Error(err))
			ctx.JSON(400, "Failed bind json")
			return
		}

		err = h.service.UpdateUser(contx, userID, p)
		if err != nil {
			logger.GetLoggerFromCtx(contx).Info(contx, "Failed update user", zap.Error(err))
			ctx.JSON(500, "Failed update user")
			return
		}

		ctx.JSON(200, "Successfully updated user")
	}
}

// DeleteUser godoc
// @Summary Delete person
// @Description Delete person by ID
// @Tags Persons
// @Accept json
// @Produce json
// @Param user_id path int true "Person ID"
// @Success 200 {object} string
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /api/users/delete/{user_id} [delete]
func (h Handler) DeleteUser(contx context.Context) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userID, err := strconv.Atoi(ctx.Param("user_id"))
		if err != nil {
			logger.GetLoggerFromCtx(contx).Info(contx, "Failed param user_id", zap.Error(err))
			ctx.JSON(400, "Bad request: Failed param user_id")
			return
		}

		err = h.service.DeleteUser(contx, userID)
		if err != nil {
			logger.GetLoggerFromCtx(contx).Info(contx, "Failed update user", zap.Error(err))
			ctx.JSON(500, "Failed update user")
			return
		}

		ctx.JSON(200, "Succesfully deleted user")
	}
}

// FindWithFilter godoc
// @Summary Find persons with filter
// @Description Get list of persons with filtering options
// @Tags Persons
// @Accept json
// @Produce json
// @Param input body model.Filter true "Filter criteria"
// @Success 200 {array} model.Person
// @Failure 204 {object} string
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /api/users/find [post]
func (h Handler) FindWithFilter(contx context.Context) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var filter model.Filter
		if err := ctx.ShouldBindJSON(&filter); err != nil {
			logger.GetLoggerFromCtx(contx).Info(contx, "Failed bing json", zap.Error(err))
			ctx.JSON(400, "Bad request")
			return
		}

		persons, err := h.service.FindWithFilter(ctx, filter)
		if err == pgx.ErrNoRows {
			ctx.JSON(204, "No content")
			return
		}
		if err != nil {
			logger.GetLoggerFromCtx(contx).Info(contx, "Failed get users using filter", zap.Error(err))
			ctx.JSON(500, "Failed get users using filter")
			return
		}

		ctx.JSON(200, persons)
	}
}
