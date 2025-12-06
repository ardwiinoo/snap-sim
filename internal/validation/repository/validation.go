package repository

import (
	"context"
	"errors"

	"github.com/ardwiinoo/snap-sim/internal/validation/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	ErrClientNotFound = errors.New("client not found")
)

type ValidationRepository struct {
	db *pgxpool.Pool
}

func NewValidationRepository(db *pgxpool.Pool) *ValidationRepository {
	return &ValidationRepository{db: db}
}

func (r *ValidationRepository) GetByClientKey(ctx context.Context, clientKey string) (*model.Client, error) {
	row := r.db.QueryRow(ctx,
		`SELECT id, client_key, client_secret, callback_url, status
         FROM snap.clients WHERE client_key = $1`,
		clientKey,
	)

	var client model.Client
	err := row.Scan(
		&client.ID,
		&client.ClientKey,
		&client.ClientSecret,
		&client.CallbackURL,
		&client.Status,
	)

	if err != nil {
		return nil, ErrClientNotFound
	}

	return &client, nil
}
