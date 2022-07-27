package service

import (
	"cloud.google.com/go/cloudsqlconn"
	"context"
	"database/sql"
	"fmt"
	"github.com/dtc03012/me/db/repository"
	"github.com/dtc03012/me/db/repository/mocks"
	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"net"
	"os"
)

const (
	EnvProd = "prod"
	EnvTest = "test"
)

//go:generate mockery --name DBService --case underscore --output ./mocks
type DBService interface {
	BeginTx(ctx context.Context, opts *sql.TxOptions) (*sqlx.Tx, error)

	CheckAdminPassword(ctx context.Context, tx *sqlx.Tx, password string) (bool, error)
	InsertAdminUUID(ctx context.Context, tx *sqlx.Tx, uuid string) error
	FindAdminUUID(ctx context.Context, tx *sqlx.Tx, uuid string) (bool, error)
}

type dbService struct {
	AdminRepo repository.Admin
}

func (dbs *dbService) BeginTx(ctx context.Context, opts *sql.TxOptions) (*sqlx.Tx, error) {

	var (
		password               = os.Getenv("MYSQL_PASSWORD")
		mysqlIP                = os.Getenv("MYSQL_IP")
		instanceConnectionName = os.Getenv("INSTANCE_CONNECTION_NAME")
		usePrivate             = os.Getenv("PRIVATE_IP")
		dataSourceName         string
	)

	d, err := cloudsqlconn.NewDialer(ctx)
	if err != nil {
		return nil, fmt.Errorf("cloudsqlconn.NewDialer: %v", err)
	}

	mysql.RegisterDialContext("cloudsqlconn",
		func(ctx context.Context, addr string) (net.Conn, error) {
			if usePrivate != "" {
				return d.Dial(ctx, instanceConnectionName, cloudsqlconn.WithPrivateIP())
			}
			return d.Dial(ctx, instanceConnectionName)
		})

	if os.Getenv("ME_ENV") == EnvProd {
		dataSourceName = fmt.Sprintf("root:%s@tcp(%s)/me?multiStatements=true", password, mysqlIP)
	} else {
		dataSourceName = fmt.Sprintf("root:%s@tcp(%s)/me_test?multiStatements=true", password, mysqlIP)
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
