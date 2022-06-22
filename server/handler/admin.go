package handler

import (
	"context"
	"github.com/dtc03012/me/protobuf/proto/service/message"
)

func (m *MeServer) CheckAdminPassword(ctx context.Context, req *message.CheckAdminPasswordRequest) (*message.CheckAdminPasswordResponse, error) {
	tx, err := m.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	defer tx.Rollback()

	isCorrect, err := m.db.CheckAdminPassword(ctx, tx, req.Password)
	if err != nil {
		return nil, err
	}

	return &message.CheckAdminPasswordResponse{
		IsAdmin: isCorrect,
	}, err
}

func (m *MeServer) FindAdminUUID(ctx context.Context, req *message.FindAdminUUIDRequest) (*message.FindAdminUUIDResponse, error) {
	tx, err := m.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	defer tx.Rollback()

	isFind, err := m.db.FindAdminUUID(ctx, tx, req.Uuid)
	if err != nil {
		return nil, err
	}

	return &message.FindAdminUUIDResponse{
		IsFind: isFind,
	}, nil
}

func (m *MeServer) InsertAdminUUID(ctx context.Context, req *message.InsertAdminUUIDRequest) (*message.InsertAdminUUIDResponse, error) {
	tx, err := m.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	defer tx.Rollback()

	err = m.db.InsertAdminUUID(ctx, tx, req.Uuid)
	if err != nil {
		return nil, err
	}

	return &message.InsertAdminUUIDResponse{}, nil
}
