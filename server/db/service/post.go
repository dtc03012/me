package service

import (
	"context"
	"errors"
	"github.com/dtc03012/me/db/entity"
	"github.com/dtc03012/me/db/option"
	"github.com/dtc03012/me/protobuf/proto/entity/post"
	"github.com/jmoiron/sqlx"
)

func (dbs *dbService) UploadPost(ctx context.Context, tx *sqlx.Tx, postData *post.Data) error {

	if postData == nil {
		return errors.New("upload post db service error: postData is nil")
	}

	p := &entity.Post{
		Title:            postData.GetTitle(),
		Password:         postData.GetPassword(),
		Writer:           postData.GetWriter(),
		Content:          postData.GetContent(),
		IsNotice:         postData.GetIsNotice(),
		Likes:            postData.GetLikes(),
		TimeToReadMinute: postData.GetTimeToReadMinute(),
	}

	tags := make([]string, 0, len(postData.GetTags()))
	for _, tag := range postData.GetTags() {
		tags = append(tags, tag)
	}

	pid, err := dbs.PostRepo.InsertPost(ctx, tx, p)
	if err != nil {
		return err
	}

	err = dbs.PostRepo.InsertBulkTag(ctx, tx, int32(pid), tags)

	return err
}

func (dbs *dbService) FetchPostList(ctx context.Context, tx *sqlx.Tx, opt *option.PostOption) ([]*post.Data, error) {

	var (
		validPostList []*entity.Post
		postList      []*entity.Post
		err           error
	)

	if opt == nil {
		return nil, errors.New("query post list db service error: option is nil")
	}

	if opt.SizeRange.Size <= 0 || opt.SizeRange.Row <= 0 {
		return nil, errors.New("fetch posts db service error: row or size is out of range")
	}

	postList, err = dbs.PostRepo.GetBulkPost(ctx, tx, opt)
	if err != nil {
		return nil, err
	}

	for i, _ := range postList {
		postList[i].Tags, err = dbs.PostRepo.GetBulkTag(ctx, tx, postList[i].Id)
		if err != nil {
			return nil, err
		}

		suc := true
		for _, tag := range opt.Tags {
			check := false
			for _, postTag := range postList[i].Tags {
				if postTag == tag {
					check = true
				}
			}

			if check == false {
				suc = false
			}
		}

		if suc == true {
			validPostList = append(validPostList, postList[i])
		}
	}

	convertPostList := convertEntityPostList(validPostList)

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

	p.Tags, err = dbs.PostRepo.GetBulkTag(ctx, tx, int32(postId))
	if err != nil {
		return nil, err
	}

	convertPost := convertEntityPost(p)

	return convertPost, nil
}

func (dbs *dbService) UpdatePost(ctx context.Context, tx *sqlx.Tx, postData *post.Data) error {

	if postData == nil {
		return errors.New("update post db service error: post is nil")
	}

	p := &entity.Post{
		Id:               postData.GetId(),
		Title:            postData.GetTitle(),
		Writer:           postData.GetWriter(),
		Content:          postData.GetContent(),
		IsNotice:         postData.GetIsNotice(),
		TimeToReadMinute: postData.GetTimeToReadMinute(),
	}

	err := dbs.PostRepo.UpdatePost(ctx, tx, p)
	if err != nil {
		return err
	}

	err = dbs.PostRepo.DeleteBulkTag(ctx, tx, postData.Id)
	if err != nil {
		return err
	}

	err = dbs.PostRepo.InsertBulkTag(ctx, tx, postData.Id, postData.Tags)
	return err
}

func (dbs *dbService) DeletePost(ctx context.Context, tx *sqlx.Tx, postId int) error {

	if postId <= 0 {
		return errors.New("delete post db service error: postId is out of range")
	}

	err := dbs.PostRepo.DeletePost(ctx, tx, int32(postId))

	return err
}

func (dbs *dbService) GetTotalPostCount(ctx context.Context, tx *sqlx.Tx, opt *option.PostOption) (int32, error) {

	totalCount, err := dbs.PostRepo.GetTotalPostCount(ctx, tx, opt)
	if err != nil {
		return 0, err
	}

	return int32(totalCount), nil
}

func (dbs *dbService) CheckPostPassword(ctx context.Context, tx *sqlx.Tx, pid int, password string) (bool, error) {

	if pid <= 0 {
		return false, errors.New("check post password db service error: pid is out of range")
	}

	check, err := dbs.PostRepo.CheckPostPassword(ctx, tx, int32(pid), password)
	if err != nil {
		return false, err
	}

	return check, err
}

func (dbs *dbService) IncrementViews(ctx context.Context, tx *sqlx.Tx, postId int, uuid string) error {

	if postId <= 0 || len(uuid) == 0 {
		return errors.New("fetch posts db service error: pid or uuid is out of range")
	}

	err := dbs.PostRepo.InsertViews(ctx, tx, int32(postId), uuid)

	return err
}

func (dbs *dbService) LeaveComment(ctx context.Context, tx *sqlx.Tx, comment *post.Comment) error {

	if comment == nil {
		return errors.New("leave comment db service error: comment is nil")
	}

	convertComment := &entity.Comment{
		PostId:          comment.PostId,
		Writer:          comment.Writer,
		ParentCommentId: comment.ParentCid,
		IsExist:         comment.IsExist,
		Password:        comment.Comment,
		Comment:         comment.Comment,
		LikeCnt:         comment.LikeCnt,
	}

	err := dbs.PostRepo.InsertComment(ctx, tx, convertComment)

	return err
}

func (dbs *dbService) FetchCommentList(ctx context.Context, tx *sqlx.Tx, opt *option.CommentOption) ([]*post.Comment, error) {

	if opt == nil {
		return nil, errors.New("leave comment db service error: option is nil")
	}

	commentList, err := dbs.PostRepo.GetBulkComment(ctx, tx, opt)
	if err != nil {
		return nil, err
	}

	convertCommentList := convertEntityCommentList(commentList)

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

func (dbs *dbService) GetTotalCommentCount(ctx context.Context, tx *sqlx.Tx, pid int) (int32, error) {

	if pid <= 0 {
		return 0, errors.New("get total comment count db service error: pid is out of range")
	}

	totalCount, err := dbs.PostRepo.GetTotalCommentCount(ctx, tx, int32(pid))
	if err != nil {
		return 0, err
	}

	return int32(totalCount), nil
}

func (dbs *dbService) CheckUserLike(ctx context.Context, tx *sqlx.Tx, pid int, uuid string) (bool, error) {

	if pid <= 0 {
		return false, errors.New("increment like db service error: pid or uuid is out of range")
	}

	check, err := dbs.PostRepo.CheckUserLike(ctx, tx, int32(pid), uuid)
	return check, err
}

func (dbs *dbService) IncrementLike(ctx context.Context, tx *sqlx.Tx, pid int, uuid string) error {

	if pid <= 0 || len(uuid) == 0 {
		return errors.New("increment like db service error: pid or uuid is out of range")
	}

	err := dbs.PostRepo.InsertLike(ctx, tx, int32(pid), uuid)
	return err
}

func (dbs *dbService) DecrementLike(ctx context.Context, tx *sqlx.Tx, pid int, uuid string) error {

	if pid <= 0 || len(uuid) == 0 {
		return errors.New("decrement like db service error: pid or uuid is out of range")
	}

	err := dbs.PostRepo.DeleteLike(ctx, tx, int32(pid), uuid)
	return err
}
