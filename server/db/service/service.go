package service

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/dtc03012/me/db/repository"
	"github.com/dtc03012/me/db/repository/mocks"
	"github.com/jmoiron/sqlx"
	"os"
)

//go:generate mockery --name DBService --case underscore --inpackage
type DBService interface {
	BeginTx(ctx context.Context, opts *sql.TxOptions) (*sqlx.Tx, error)

	CheckAdmin(ctx context.Context, tx *sqlx.Tx, password string) (bool, error)
}

type dbService struct {
	AdminRepo repository.Admin
}

func (dbs *dbService) BeginTx(ctx context.Context, opts *sql.TxOptions) (*sqlx.Tx, error) {

	password := os.Getenv("MYSQL_PASSWORD")
	dataSourceName := fmt.Sprintf("user:%s@tcp(host:3306)/me?multiStatements=true", password)

	db, err := sqlx.Connect("mysql", dataSourceName)
	if err != nil {
		return nil, err
	}

	tx, err := db.BeginTxx(ctx, opts)

	if err != nil {
		return nil, err
	}

	return tx, nil
}

func NewDBService() DBService {
	return &dbService{
		AdminRepo: repository.NewAdminRepo(),
	}
}

func NewMockDBService() *dbService {
	return &dbService{
		AdminRepo: &mocks.Admin{},
	}
}
