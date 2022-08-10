package service

import (
	"cloud.google.com/go/cloudsqlconn"
	"context"
	"database/sql"
	"fmt"
	"github.com/dtc03012/me/db/option"
	"github.com/dtc03012/me/db/repository"
	"github.com/dtc03012/me/db/repository/mocks"
	"github.com/dtc03012/me/protobuf/proto/entity/post"
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

	CheckAdminPassword(ctx context.Context, tx *sqlx.Tx, password string) (bool, error)
	InsertAdminUUID(ctx context.Context, tx *sqlx.Tx, uuid string) error
	FindAdminUUID(ctx context.Context, tx *sqlx.Tx, uuid string) (bool, error)

	UploadPost(ctx context.Context, tx *sqlx.Tx, postData *post.Data) error
	FetchPostList(ctx context.Context, tx *sqlx.Tx, row int, size int) ([]*post.Data, error)
	FetchPost(ctx context.Context, tx *sqlx.Tx, postId int) (*post.Data, error)
	IncrementViews(ctx context.Context, tx *sqlx.Tx, postId int, uuid string) error

	LeaveComment(ctx context.Context, tx *sqlx.Tx, comment *post.Comment) error
	FetchCommentList(ctx context.Context, tx *sqlx.Tx, opt *option.CommentOption) ([]*post.Comment, error)
	DeleteComment(ctx context.Context, tx *sqlx.Tx, postId int, commentId int) error

	QueryPostList(ctx context.Context, tx *sqlx.Tx, opt *option.PostOption) ([]*post.Data, error)
}

type dbService struct {
	AdminRepo repository.Admin
	PostRepo  repository.Post
}

func (dbs *dbService) BeginTx(ctx context.Context, opts *sql.TxOptions) (*sqlx.Tx, error) {

	var (
		password               = os.Getenv("MYSQL_PASSWORD")
		mysqlIP                = os.Getenv("MYSQL_IP")
		instanceConnectionName = os.Getenv("INSTANCE_CONNECTION_NAME")
		usePrivate             = os.Getenv("PRIVATE_IP")
		env                    = os.Getenv("ME_ENV")
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
		dataSourceName = fmt.Sprintf("root:%s@tcp(%s)/me?multiStatements=true&parseTime=true", password, mysqlIP)
	} else {
		dataSourceName = fmt.Sprintf("root:%s@tcp(%s)/me_test?multiStatements=true&parseTime=true", password, mysqlIP)
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
		AdminRepo: repository.NewAdminRepo(),
		PostRepo:  repository.NewPostRepo(),
	}
}

type mockDBService struct {
	AdminRepo mocks.Admin
	PostRepo  mocks.Post
}

func NewMockDBService() (DBService, *mockDBService) {
	m := mockDBService{
		AdminRepo: mocks.Admin{},
	}

	return &dbService{
		AdminRepo: &m.AdminRepo,
		PostRepo:  &m.PostRepo,
	}, &m
}
