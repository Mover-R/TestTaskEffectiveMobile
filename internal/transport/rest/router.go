package server

import (
	"TestTaskEffectiveMobile/internal/config"
	"TestTaskEffectiveMobile/pkg/logger"
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Router struct {
	Config    *config.Config
	RestServe *gin.Engine
}

func NewRouter(ctx context.Context, cfg *config.Config) (Router, error) {
	r := gin.Default()
	urlFront := cfg.FrontendURL
	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", urlFront)
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Methods", "GET, POST, OPTIONS, DELETE")
		c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Accept, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204) // No Content
			return
		}

		c.Next()
	})
	r.Use(logger.MiddleWare(ctx, logger.GetLoggerFromCtx(ctx)))
	r.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "Its OK"})
	})
	return Router{Config: cfg, RestServe: r}, nil
}

func (r *Router) Run(ctx context.Context) {
	linq := fmt.Sprintf("%s:%s", r.Config.RestHost, r.Config.RestPort)
	go func() {
		if err := r.RestServe.Run(linq); err != nil {
			logger.GetLoggerFromCtx(ctx).Fatal(ctx, "Failed to serve:", zap.Error(err))
		}
	}()
	logger.GetLoggerFromCtx(ctx).Info(ctx, "Succesfully Served")
}
