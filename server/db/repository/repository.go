package repository

import (
	"context"
	"github.com/dtc03012/me/db/entity"
	"github.com/dtc03012/me/db/option"
	"github.com/jmoiron/sqlx"
)

//go:generate mockery --name Admin --case underscore --output ./mocks
type Admin interface {
	GetPassword(ctx context.Context, tx *sqlx.Tx) (string, error)
	InsertUUID(ctx context.Context, tx *sqlx.Tx, uuid string) error
	FindUUID(ctx context.Context, tx *sqlx.Tx, uuid string) (string, error)
}

//go:generate mockery --name Post --case underscore --output ./mocks
type Post interface {
	GetPost(ctx context.Context, tx *sqlx.Tx, pid int32) (*entity.Post, error)
	GetBulkPost(ctx context.Context, tx *sqlx.Tx, opt *option.PostOption) ([]*entity.Post, error)
	InsertPost(ctx context.Context, tx *sqlx.Tx, post *entity.Post) (int, error)
	DeletePost(ctx context.Context, tx *sqlx.Tx, pid int32) error
	UpdatePost(ctx context.Context, tx *sqlx.Tx, post *entity.Post) error
	GetTotalPostCount(ctx context.Context, tx *sqlx.Tx, opt *option.PostOption) (int, error)
	CheckPostPassword(ctx context.Context, tx *sqlx.Tx, pid int32, password string) (bool, error)

	GetBulkTag(ctx context.Context, tx *sqlx.Tx, pid int32) ([]string, error)
	InsertBulkTag(ctx context.Context, tx *sqlx.Tx, pid int32, tags []string) error
	DeleteBulkTag(ctx context.Context, tx *sqlx.Tx, pid int32) error

	GetViews(ctx context.Context, tx *sqlx.Tx, pid int32) (int, error)
	InsertViews(ctx context.Context, tx *sqlx.Tx, pid int32, uuid string) error

	GetComment(ctx context.Context, tx *sqlx.Tx, cid int32) ([]*entity.Comment, error)
	GetBulkComment(ctx context.Context, tx *sqlx.Tx, opt *option.CommentOption) ([]*entity.Comment, error)
	InsertComment(ctx context.Context, tx *sqlx.Tx, comment *entity.Comment) error
	DeleteComment(ctx context.Context, tx *sqlx.Tx, commentId int, password string) error
	GetTotalCommentCount(ctx context.Context, tx *sqlx.Tx, pid int32) (int, error)
	CheckCommentPassword(ctx context.Context, tx *sqlx.Tx, cid int32, password string) (bool, error)

	CheckUserLike(ctx context.Context, tx *sqlx.Tx, pid int32, uuid string) (bool, error)
	InsertLike(ctx context.Context, tx *sqlx.Tx, pid int32, uuid string) error
	DeleteLike(ctx context.Context, tx *sqlx.Tx, pid int32, uuid string) error
}

func NewAdminRepo() Admin {
	return &admin{}
}

func NewPostRepo() Post {
	return &post{}
}
