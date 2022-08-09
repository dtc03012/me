package service

import (
	"cloud.google.com/go/cloudsqlconn"
	"context"
	"database/sql"
	"fmt"
	"github.com/dtc03012/me/file_server/db/repository"
	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"net"
	"os"
)

const (
	EnvProd = "prod"
	EnvTest = "test"
	EnvDev  = "dev"
)

//go:generate mockery --name DBService --case underscore --output ./mocks
type DBService interface {
	BeginTx(ctx context.Context, opts *sql.TxOptions) (*sqlx.Tx, error)

	GetFileId(ctx context.Context, tx *sqlx.Tx, fileName string) (string, error)
	InsertFileInfo(ctx context.Context, tx *sqlx.Tx, fileName string, fileId string) error
}

type dbService struct {
	FileRepo repository.File
}

func (dbs *dbService) BeginTx(ctx context.Context, opts *sql.TxOptions) (*sqlx.Tx, error) {

	var (
		password               = os.Getenv("MYSQL_PASSWORD")
		mysqlIP                = os.Getenv("MYSQL_IP")
		instanceConnectionName = os.Getenv("INSTANCE_CONNECTION_NAME")
		usePrivate             = os.Getenv("PRIVATE_IP")
		env                    = os.Getenv("FILE_SERVER_ENV")
		dataSourceName         string
	)

	if env == EnvProd {
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
	}

	if env == EnvProd || env == EnvDev {
		dataSourceName = fmt.Sprintf("root:%s@tcp(%s)/file_server?multiStatements=true&parseTime=true", password, mysqlIP)
	} else {
		dataSourceName = fmt.Sprintf("root:%s@tcp(%s)/file_server_test?multiStatements=true&parseTime=true", password, mysqlIP)
	}

	db, err := sqlx.Connect("mysql", dataSourceName)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	tx, err := db.BeginTxx(ctx, opts)

	if err != nil {
		return nil, err
	}

	return tx, nil
}

func NewDBService() DBService {
	return &dbService{
		FileRepo: repository.NewFileRepo(),
	}
}

//type mockDBService struct {
//	AdminRepo mocks.Admin
//	PostRepo  mocks.Post
//}
//
//func NewMockDBService() (DBService, *mockDBService) {
//	m := mockDBService{
//		AdminRepo: mocks.Admin{},
//	}
//
//	return &dbService{
//		AdminRepo: &m.AdminRepo,
//		PostRepo:  &m.PostRepo,
//	}, &m
//}
