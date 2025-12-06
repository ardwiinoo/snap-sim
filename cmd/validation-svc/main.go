package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/ardwiinoo/snap-sim/internal/common/config"
	"github.com/ardwiinoo/snap-sim/internal/common/db"
	"github.com/ardwiinoo/snap-sim/internal/common/logger"
	"github.com/ardwiinoo/snap-sim/internal/validation/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.Load("configs/validation.yaml")
	if err != nil {
		panic(fmt.Errorf("failed to load config: %w", err))
	}

	log := logger.New()

	pool, err := db.NewPgPool(cfg.DB)
	if err != nil {
		log.Error("failed to connect to db", slog.Any("err", err))
		os.Exit(1)
	}

	r := gin.Default()

	validationHandler := handler.New(pool, log)
	r.POST("/v1/payment/validate", validationHandler.Validate)

	addr := fmt.Sprintf(":%d", cfg.Server.Port)
	log.Info("starting validation service", slog.String("port", addr))

	if err := r.Run(addr); err != nil {
		log.Error("server failed to start", slog.Any("err", err))
	}
}
