package service

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/dtc03012/me/db/repository"
	"github.com/dtc03012/me/db/repository/mocks"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"os"
)

const (
	EnvProd = "prod"
	EnvTest = "test"
)

//go:generate mockery --name DBService --case underscore --output ./mocks
type DBService interface {
	BeginTx(ctx context.Context, opts *sql.TxOptions) (*sqlx.Tx, error)

	CheckAdmin(ctx context.Context, tx *sqlx.Tx, password string) (bool, error)
}

type dbService struct {
	AdminRepo repository.Admin
}

func (dbs *dbService) BeginTx(ctx context.Context, opts *sql.TxOptions) (*sqlx.Tx, error) {

	var (
		password       string
		dataSourceName string
	)
	password = os.Getenv("MYSQL_PASSWORD")

	if os.Getenv("GOENV") == EnvProd {
		dataSourceName = fmt.Sprintf("root:%s@tcp(localhost:3306)/me?multiStatements=true", password)
	} else {
		dataSourceName = fmt.Sprintf("root:%s@tcp(localhost:3306)/me_test?multiStatements=true", password)
	}

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

type mockDBService struct {
	AdminRepo mocks.Admin
}

func NewMockDBService() (DBService, *mockDBService) {
	m := mockDBService{
		AdminRepo: mocks.Admin{},
	}

	return &dbService{
		AdminRepo: &m.AdminRepo,
	}, &m
}
