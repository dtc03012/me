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
		return errors.New("upload post error: postData is nil")
	}

	post := &entity.Post{
		Title:            postData.GetTitle(),
		Writer:           postData.GetWriter(),
		Content:          postData.GetContent(),
		TimeToReadMinute: postData.GetTimeToReadMinute(),
	}

	tags := make([]string, 0, len(postData.GetTags()))
	for _, tag := range postData.GetTags() {
		tags = append(tags, tag)
	}

	err := dbs.PostRepo.InsertPost(ctx, tx, post, tags)

	return err
}

func (dbs *dbService) FetchPosts(ctx context.Context, tx *sqlx.Tx, row int, size int) ([]*post.Data, error) {

	if row <= 0 || size <= 0 {
		return nil, errors.New("fetch posts error: row or size is out of range")
	}

	posts, err := dbs.PostRepo.GetBulkPost(ctx, tx, &option.PostOption{
		SizeRange: &option.RangeOption{
			Row:  int32(row),
			Size: int32(size),
		},
	})

	if err != nil {
		return nil, err
	}

	convertPosts := make([]*post.Data, 0, len(posts))

	for _, p := range posts {
		convertPost := &post.Data{
			Id:               p.Id,
			Title:            p.Title,
			Writer:           p.Writer,
			Content:          p.Content,
			Tags:             p.Tags,
			TimeToReadMinute: p.TimeToReadMinute,
			CreateAt:         timestamppb.New(p.CreateAt),
		}

		convertPosts = append(convertPosts, convertPost)
	}

	return convertPosts, nil
}
