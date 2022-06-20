package repository

import (
	"context"
	"github.com/jmoiron/sqlx"
)

type Admin interface {
	GetPassword(ctx context.Context, tx *sqlx.Tx) (string, error)
}

func NewAdminRepo() Admin {
	return &admin{}
}
