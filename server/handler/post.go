package handler

import (
	"context"
	"github.com/dtc03012/me/db/option"
	"github.com/dtc03012/me/protobuf/proto/service/message"
)

func (m *MeServer) UploadPost(ctx context.Context, req *message.UploadPostRequest) (*message.UploadPostResponse, error) {

	tx, err := m.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	defer tx.Rollback()

	err = m.db.UploadPost(ctx, tx, req.GetData())
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

	posts, err := m.db.FetchPostList(ctx, tx, &option.PostOption{
		SizeRange: &option.RangeOption{
			Row:  int(req.GetRow()),
			Size: int(req.GetSize()),
		},
		ClassificationType: option.ClassificationTypeMap[req.GetOption().GetClassificationOption().String()],
		QueryType:          option.QueryTypeMap[req.GetOption().GetQueryOption().String()],
		Query:              req.GetOption().GetQuery(),
	})
	if err != nil {
		return nil, err
	}

	totalPostCount, err := m.db.GetTotalPostCount(ctx, tx)

	tx.Commit()
	return &message.FetchPostListResponse{PostList: posts, TotalPostCount: totalPostCount}, nil
}

func (m *MeServer) FetchPost(ctx context.Context, req *message.FetchPostRequest) (*message.FetchPostResponse, error) {

	tx, err := m.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	defer tx.Rollback()

	post, err := m.db.FetchPost(ctx, tx, int(req.Id))

	if err != nil {
		return nil, err
	}

	tx.Commit()
	return &message.FetchPostResponse{Post: post}, nil
}

func (m *MeServer) IncrementView(ctx context.Context, req *message.IncrementViewRequest) (*message.IncrementViewResponse, error) {

	tx, err := m.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	defer tx.Rollback()

	err = m.db.IncrementViews(ctx, tx, int(req.GetId()), req.GetUuid())
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

	err = m.db.DeleteComment(ctx, tx, int(req.GetPostId()), int(req.GetCommentId()))
	if err != nil {
		return nil, err
	}

	tx.Commit()
	return &message.DeleteCommentResponse{}, nil
}
