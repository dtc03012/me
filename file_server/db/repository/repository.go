package repository

import (
	"context"
	"github.com/jmoiron/sqlx"
)

//go:generate mockery --name Post --case underscore --output ./mocks
type File interface {
	GetFileId(ctx context.Context, tx *sqlx.Tx, fileName string) (string, error)
	InsertFileIdName(ctx context.Context, tx *sqlx.Tx, fileName string, fileId string) error
}

func NewFileRepo() File {
	return &file{}
}
