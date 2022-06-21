package db

import (
	"context"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
)

func SetupMock() (context.Context, *sqlx.Tx, sqlmock.Sqlmock, error) {
	db, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, nil, err
	}

	sqlxDB := sqlx.NewDb(db, "sqlmock")

	ctx := context.Background()

	mock.ExpectBegin()
	tx, err := sqlxDB.BeginTxx(ctx, nil)

	if err != nil {
		return nil, nil, nil, err
	}

	return ctx, tx, mock, err
}
