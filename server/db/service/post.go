package service

import (
	"context"
	"errors"
	"github.com/dtc03012/me/db/entity"
	"github.com/dtc03012/me/db/option"
	"github.com/dtc03012/me/protobuf/proto/entity/post"
	"github.com/jmoiron/sqlx"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (dbs *dbService) UploadPost(ctx context.Context, tx *sqlx.Tx, postData *post.Data) error {

	if postData == nil {
		return errors.New("upload post db service error: postData is nil")
	}

	post := &entity.Post{
		Title:            postData.GetTitle(),
		Writer:           postData.GetWriter(),
		Content:          postData.GetContent(),
		LikeCnt:          postData.GetLikeCnt(),
		TimeToReadMinute: postData.GetTimeToReadMinute(),
	}

	tags := make([]string, 0, len(postData.GetTags()))
	for _, tag := range postData.GetTags() {
		tags = append(tags, tag)
	}

	err := dbs.PostRepo.InsertPost(ctx, tx, post, tags)

	return err
}

func (dbs *dbService) FetchPostList(ctx context.Context, tx *sqlx.Tx, row int, size int) ([]*post.Data, error) {

	if row <= 0 || size <= 0 {
		return nil, errors.New("fetch posts db service error: row or size is out of range")
	}

	postList, err := dbs.PostRepo.GetBulkPost(ctx, tx, &option.PostOption{
		SizeRange: &option.RangeOption{
			Row:  int32(row),
			Size: int32(size),
		},
	})

	if err != nil {
		return nil, err
	}

	convertPostList := make([]*post.Data, 0, len(postList))

	for _, p := range postList {
		convertPost := &post.Data{
			Id:               p.Id,
			Title:            p.Title,
			Writer:           p.Writer,
			Content:          p.Content,
			Tags:             p.Tags,
			Views:            p.Views,
			TimeToReadMinute: p.TimeToReadMinute,
			LikeCnt:          p.LikeCnt,
			CreateAt:         timestamppb.New(p.CreateAt),
		}

		convertPostList = append(convertPostList, convertPost)
	}

	return convertPostList, nil
}

func (dbs *dbService) FetchPost(ctx context.Context, tx *sqlx.Tx, postId int) (*post.Data, error) {

	if postId <= 0 {
		return nil, errors.New("fetch posts db service error: postId is out of range")
	}

	p, err := dbs.PostRepo.GetPost(ctx, tx, int32(postId))

	if err != nil {
		return nil, err
	}

	convertPost := &post.Data{
		Id:               p.Id,
		Title:            p.Title,
		Writer:           p.Writer,
		Content:          p.Content,
		Tags:             p.Tags,
		Views:            p.Views,
		TimeToReadMinute: p.TimeToReadMinute,
		LikeCnt:          p.LikeCnt,
		CreateAt:         timestamppb.New(p.CreateAt),
	}

	return convertPost, nil
}

func (dbs *dbService) IncrementViews(ctx context.Context, tx *sqlx.Tx, postId int) error {

	if postId <= 0 {
		return errors.New("fetch posts db service error: pid is out of range")
	}

	views, err := dbs.PostRepo.GetViews(ctx, tx, int32(postId))
	if err != nil {
		return err
	}

	err = dbs.PostRepo.UpdateViews(ctx, tx, int32(views+1), int32(postId))

	return err
}

func (dbs *dbService) LeaveComment(ctx context.Context, tx *sqlx.Tx, comment *post.Comment) error {

	if comment == nil {
		return errors.New("leave comment db service error: comment is nil")
	}

	convertComment := &entity.Comment{
		PostId:   comment.PostId,
		Writer:   comment.Writer,
		Password: comment.Comment,
		Comment:  comment.Comment,
		LikeCnt:  comment.LikeCnt,
	}

	err := dbs.PostRepo.InsertComment(ctx, tx, convertComment)

	return err
}

func (dbs *dbService) FetchCommentList(ctx context.Context, tx *sqlx.Tx, postId int, row int, size int) ([]*post.Comment, error) {

	if postId <= 0 {
		return nil, errors.New("leave comment db service error: comment is nil")
	}

	commentList, err := dbs.PostRepo.GetBulkComment(ctx, tx, &option.CommentOption{
		SizeRange: &option.RangeOption{
			Row:  int32(row),
			Size: int32(size),
		},
		PostId: postId,
	})
	if err != nil {
		return nil, err
	}

	convertCommentList := make([]*post.Comment, 0, len(commentList))
	for _, c := range commentList {
		convertComment := &post.Comment{
			Id:       c.Id,
			PostId:   c.PostId,
			Writer:   c.Writer,
			Password: c.Password,
			Comment:  c.Comment,
			LikeCnt:  c.LikeCnt,
			CreateAt: timestamppb.New(c.CreateAt),
		}

		convertCommentList = append(convertCommentList, convertComment)
	}

	return convertCommentList, nil
}

func (dbs *dbService) DeleteComment(ctx context.Context, tx *sqlx.Tx, postId int, commentId int) error {

	if postId <= 0 {
		return errors.New("delete comment db service error: post id is out of range")
	}

	if commentId <= 0 {
		return errors.New("delete comment db service error: comment id is out of range")
	}

	err := dbs.PostRepo.DeleteComment(ctx, tx, postId, commentId)

	return err
}
