package repository

import (
	"context"
	"github.com/jmoiron/sqlx"
)

//go:generate mockery --name Admin --case underscore --output ./mocks
type Admin interface {
	GetPassword(ctx context.Context, tx *sqlx.Tx) (string, error)
	InsertUUID(ctx context.Context, tx *sqlx.Tx, uuid string) error
	FindUUID(ctx context.Context, tx *sqlx.Tx, uuid string) (string, error)
}

func NewAdminRepo() Admin {
	return &admin{}
}
