package handler

import (
	"context"
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

func (m *MeServer) FetchPosts(ctx context.Context, req *message.FetchPostsRequest) (*message.FetchPostsResponse, error) {

	tx, err := m.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	defer tx.Rollback()

	posts, err := m.db.FetchPosts(ctx, tx, int(req.Row), int(req.Size))

	if err != nil {
		return nil, err
	}

	return &message.FetchPostsResponse{Data: posts}, nil
}
