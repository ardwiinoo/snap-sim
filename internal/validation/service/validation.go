package service

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"log/slog"
	"time"

	"github.com/ardwiinoo/snap-sim/internal/validation/model"
	"github.com/ardwiinoo/snap-sim/internal/validation/repository"
)

var (
	ErrClientNotFound    = errors.New("client not found")
	ErrSignatureMismatch = errors.New("signature mismatch")
	ErrTimestampInvalid  = errors.New("timestamp invalid or expired")
)

type ValidationService struct {
	repo *repository.ValidationRepository
	log  *slog.Logger
}

func New(repo *repository.ValidationRepository, log *slog.Logger) *ValidationService {
	return &ValidationService{
		repo: repo,
		log:  log,
	}
}

func (s *ValidationService) LookupClient(ctx context.Context, clientKey string) (*model.Client, error) {
	client, err := s.repo.GetByClientKey(ctx, clientKey)
	if err != nil {
		return nil, ErrClientNotFound
	}
	return client, nil
}

func (s *ValidationService) ValidateTimestamp(ts string) error {
	timestamp, err := time.Parse(time.RFC3339, ts)
	if err != nil {
		return ErrTimestampInvalid
	}

	// 5 minutes window
	now := time.Now().UTC()
	diff := now.Sub(timestamp)

	if diff > 5*time.Minute || diff < -5*time.Minute {
		return ErrTimestampInvalid
	}
	return nil
}

func (s *ValidationService) BuildStringToSign(method, path, timestamp string, body []byte) string {
	hash := sha256.Sum256(body)
	bodyHash := hex.EncodeToString(hash[:])
	return fmt.Sprintf("%s:%s:%s:%s", method, path, timestamp, bodyHash)
}

func (s *ValidationService) ValidateSignature(secret, stringToSign, given string) error {
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(stringToSign))
	expected := base64.StdEncoding.EncodeToString(mac.Sum(nil))

	if !hmac.Equal([]byte(expected), []byte(given)) {
		return ErrSignatureMismatch
	}
	return nil
}
