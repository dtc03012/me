package handler

import (
	"context"
	"errors"
	"github.com/dtc03012/me/db/option"
	"github.com/dtc03012/me/protobuf/proto/service/message"
)

func (m *MeServer) UploadPost(ctx context.Context, req *message.UploadPostRequest) (*message.UploadPostResponse, error) {

	tx, err := m.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	defer tx.Rollback()

	err = m.db.UploadPost(ctx, tx, req.GetPost())
	if err != nil {
		return nil, err
	}

	tx.Commit()
	return &message.UploadPostResponse{}, nil
}

func (m *MeServer) FetchPostList(ctx context.Context, req *message.FetchPostListRequest) (*message.FetchPostListResponse, error) {

	tx, err := m.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	defer tx.Rollback()

	postOption := &option.PostOption{
		SizeRange: &option.RangeOption{
			Row:  int(req.GetRow()),
			Size: int(req.GetSize()),
		},
		ClassificationType: option.ClassificationTypeMap[req.GetOption().GetClassificationOption().String()],
		QueryType:          option.QueryTypeMap[req.GetOption().GetQueryOption().String()],
		Query:              req.GetOption().GetQuery(),
	}

	postList, err := m.db.FetchPostList(ctx, tx, postOption)

	if err != nil {
		return nil, err
	}

	totalPostCount, err := m.db.GetTotalPostCount(ctx, tx, postOption)

	for _, post := range postList {
		post.IsLike, err = m.db.CheckUserLike(ctx, tx, int(post.Id), req.GetUuid())
		if err != nil {
			return nil, err
		}
	}

	tx.Commit()
	return &message.FetchPostListResponse{PostList: postList, TotalPostCount: totalPostCount}, nil
}

func (m *MeServer) FetchPost(ctx context.Context, req *message.FetchPostRequest) (*message.FetchPostResponse, error) {

	tx, err := m.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	defer tx.Rollback()

	post, err := m.db.FetchPost(ctx, tx, int(req.GetPostId()))
	if err != nil {
		return nil, err
	}

	post.IsLike, err = m.db.CheckUserLike(ctx, tx, int(post.Id), req.GetUuid())
	if err != nil {
		return nil, err
	}

	tx.Commit()
	return &message.FetchPostResponse{Post: post}, nil
}

func (m *MeServer) UpdatePost(ctx context.Context, req *message.UpdatePostRequest) (*message.UpdatePostResponse, error) {

	tx, err := m.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	defer tx.Rollback()

	check, err := m.db.CheckPostPassword(ctx, tx, int(req.GetPost().GetId()), req.GetPost().GetPassword())
	if err != nil {
		return nil, err
	}

	if !check {
		return nil, errors.New("password isn't correct")
	}

	err = m.db.UpdatePost(ctx, tx, req.GetPost())
	if err != nil {
		return nil, err
	}

	tx.Commit()
	return &message.UpdatePostResponse{}, nil
}

func (m *MeServer) DeletePost(ctx context.Context, req *message.DeletePostRequest) (*message.DeletePostResponse, error) {

	tx, err := m.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	defer tx.Rollback()

	check, err := m.db.CheckPostPassword(ctx, tx, int(req.GetPostId()), req.GetPassword())
	if err != nil {
		return nil, err
	}

	if !check {
		return nil, errors.New("password isn't correct")
	}

	err = m.db.DeletePost(ctx, tx, int(req.GetPostId()))
	if err != nil {
		return nil, err
	}

	tx.Commit()
	return &message.DeletePostResponse{}, nil
}

func (m *MeServer) CheckPostPassword(ctx context.Context, req *message.CheckPostPasswordRequest) (*message.CheckPostPasswordResponse, error) {

	tx, err := m.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	defer tx.Rollback()

	check, err := m.db.CheckPostPassword(ctx, tx, int(req.GetPostId()), req.GetPassword())
	if err != nil {
		return nil, err
	}

	tx.Commit()
	return &message.CheckPostPasswordResponse{
		Success: check,
	}, nil
}

func (m *MeServer) IncrementView(ctx context.Context, req *message.IncrementViewRequest) (*message.IncrementViewResponse, error) {

	tx, err := m.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	defer tx.Rollback()

	err = m.db.IncrementViews(ctx, tx, int(req.GetPostId()), req.GetUuid())
	if err != nil {
		return nil, err
	}

	tx.Commit()
	return &message.IncrementViewResponse{}, nil
}

func (m *MeServer) LeaveComment(ctx context.Context, req *message.LeaveCommentRequest) (*message.LeaveCommentResponse, error) {

	tx, err := m.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	defer tx.Rollback()

	err = m.db.LeaveComment(ctx, tx, req.GetComment())
	if err != nil {
		return nil, err
	}

	tx.Commit()
	return &message.LeaveCommentResponse{}, nil
}

func (m *MeServer) FetchCommentList(ctx context.Context, req *message.FetchCommentListRequest) (*message.FetchCommentListResponse, error) {

	tx, err := m.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	defer tx.Rollback()

	commentList, err := m.db.FetchCommentList(ctx, tx, &option.CommentOption{
		SizeRange: &option.RangeOption{
			Row:  int(req.GetRow()),
			Size: int(req.GetSize()),
		},
		PostId: int(req.GetPostId()),
	})
	if err != nil {
		return nil, err
	}

	totalCommentCount, err := m.db.GetTotalCommentCount(ctx, tx, int(req.GetPostId()))
	if err != nil {
		return nil, err
	}

	tx.Commit()
	return &message.FetchCommentListResponse{
		CommentList:       commentList,
		TotalCommentCount: totalCommentCount,
	}, nil
}

func (m *MeServer) DeleteComment(ctx context.Context, req *message.DeleteCommentRequest) (*message.DeleteCommentResponse, error) {

	tx, err := m.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	defer tx.Rollback()

	check, err := m.db.CheckCommentPassword(ctx, tx, int(req.GetCommentId()), req.GetPassword())
	if err != nil {
		return nil, err
	}

	if !check {
		return nil, errors.New("password isn't correct")
	}

	err = m.db.DeleteComment(ctx, tx, int(req.GetCommentId()), req.GetPassword())
	if err != nil {
		return nil, err
	}

	tx.Commit()
	return &message.DeleteCommentResponse{}, nil
}

func (m *MeServer) IncrementLike(ctx context.Context, req *message.IncrementLikeRequest) (*message.IncrementLikeResponse, error) {

	tx, err := m.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	defer tx.Rollback()

	err = m.db.IncrementLike(ctx, tx, int(req.GetPostId()), req.GetUuid())
	if err != nil {
		return nil, err
	}

	tx.Commit()
	return &message.IncrementLikeResponse{}, nil
}

func (m *MeServer) DecrementLike(ctx context.Context, req *message.DecrementLikeRequest) (*message.DecrementLikeResponse, error) {

	tx, err := m.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	defer tx.Rollback()

	err = m.db.DecrementLike(ctx, tx, int(req.GetPostId()), req.GetUuid())
	if err != nil {
		return nil, err
	}

	tx.Commit()
	return &message.DecrementLikeResponse{}, nil
}

func (m *MeServer) CheckValidPostId(ctx context.Context, req *message.CheckValidPostIdRequest) (*message.CheckValidPostIdResponse, error) {

	tx, err := m.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	defer tx.Rollback()

	_, err = m.db.FetchPost(ctx, tx, int(req.GetPostId()))
	if err != nil {
		return nil, err
	}

	return &message.CheckValidPostIdResponse{}, nil
}

func (m *MeServer) CheckValidCommentId(ctx context.Context, req *message.CheckValidCommentIdRequest) (*message.CheckValidCommentIdResponse, error) {

	tx, err := m.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	defer tx.Rollback()

	_, err = m.db.FetchComment(ctx, tx, int(req.GetCid()))
	if err != nil {
		return nil, err
	}

	return &message.CheckValidCommentIdResponse{}, nil
}
