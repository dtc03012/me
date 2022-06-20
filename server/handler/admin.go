package handler

import (
	"context"
	"github.com/dtc03012/me/protobuf/proto/service/message"
)

func (m *MeServer) CheckAdmin(ctx context.Context, req *message.CheckAdminRequest) (*message.CheckAdminResponse, error) {
	tx, err := m.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	defer tx.Rollback()

	isCorrect, err := m.db.CheckAdmin(ctx, tx, req.Password)
	if err != nil {
		return nil, err
	}

	return &message.CheckAdminResponse{
		IsAdmin: isCorrect,
	}, err
}
