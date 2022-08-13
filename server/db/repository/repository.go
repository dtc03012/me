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
	InsertPost(ctx context.Context, tx *sqlx.Tx, post *entity.Post, tags []string) error
	GetTotalPostCount(ctx context.Context, tx *sqlx.Tx) (int, error)

	GetBulkTag(ctx context.Context, tx *sqlx.Tx, pid int32) ([]string, error)

	GetViews(ctx context.Context, tx *sqlx.Tx, pid int32) (int, error)
	InsertViews(ctx context.Context, tx *sqlx.Tx, pid int32, uuid string) error

	GetBulkComment(ctx context.Context, tx *sqlx.Tx, opt *option.CommentOption) ([]*entity.Comment, error)
	InsertComment(ctx context.Context, tx *sqlx.Tx, comment *entity.Comment) error
	DeleteComment(ctx context.Context, tx *sqlx.Tx, postId int, commentId int) error
	GetTotalCommentCount(ctx context.Context, tx *sqlx.Tx, pid int32) (int, error)
}

func NewAdminRepo() Admin {
	return &admin{}
}

func NewPostRepo() Post {
	return &post{}
}
