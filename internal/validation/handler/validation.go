package handler

import (
	"bytes"
	"context"
	"io"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/ardwiinoo/snap-sim/internal/validation/model"
	"github.com/ardwiinoo/snap-sim/internal/validation/repository"
	"github.com/ardwiinoo/snap-sim/internal/validation/service"

	"github.com/jackc/pgx/v5/pgxpool"
)

type ValidationHandler struct {
	service *service.ValidationService
	log     *slog.Logger
}

func New(db *pgxpool.Pool, log *slog.Logger) *ValidationHandler {
	repo := repository.NewValidationRepository(db)
	svc := service.New(repo, log)

	return &ValidationHandler{
		service: svc,
		log:     log,
	}
}

func (h *ValidationHandler) Validate(c *gin.Context) {
	ctx := context.Background()

	clientKey := c.GetHeader("X-Client-Key")
	timestamp := c.GetHeader("X-Timestamp")
	signature := c.GetHeader("X-Signature")

	if clientKey == "" || timestamp == "" || signature == "" {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Code:    "400",
			Message: "Missing authentication headers",
		})
		return
	}

	body, err := c.GetRawData()
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Code:    "400",
			Message: "Invalid request body",
		})
		return
	}

	// Restore body for downstream use
	c.Request.Body = io.NopCloser(bytes.NewBuffer(body))

	if err := h.service.ValidateTimestamp(timestamp); err != nil {
		c.JSON(http.StatusUnauthorized, model.ErrorResponse{
			Code:    "401",
			Message: "Timestamp expired or invalid",
		})
		return
	}

	client, err := h.service.LookupClient(ctx, clientKey)
	if err != nil {
		c.JSON(http.StatusUnauthorized, model.ErrorResponse{
			Code:    "401",
			Message: "Invalid client",
		})
		return
	}

	// Build string-to-sign
	stringToSign := h.service.BuildStringToSign(
		c.Request.Method,
		c.FullPath(),
		timestamp,
		body,
	)

	// Validate signature
	if err := h.service.ValidateSignature(client.ClientSecret, stringToSign, signature); err != nil {
		c.JSON(http.StatusUnauthorized, model.ErrorResponse{
			Code:    "401",
			Message: "Invalid signature",
		})
		return
	}

	// TODO: forward to payment adapter service

	c.JSON(http.StatusOK, model.SuccessResponse{
		Status:  "success",
		Message: "Valid signature",
	})
}
