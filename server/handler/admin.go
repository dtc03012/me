package handler

import (
	"context"
	"fmt"
	"github.com/dtc03012/me/protobuf/proto/service/message"
)

func (m *MeServer) LoginAdmin(ctx context.Context, req *message.LoginAdminRequest) (*message.LoginAdminResponse, error) {
	tx, err := m.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	defer tx.Rollback()

	isCorrect, err := m.db.CheckAdminPassword(ctx, tx, req.Password)
	if err != nil {
		return nil, err
	}

	if isCorrect == true {
		err := m.db.InsertAdminUUID(ctx, tx, req.Uuid)
		if err != nil {
			return nil, err
		}
	}

	tx.Commit()
	return &message.LoginAdminResponse{
		IsAdmin: isCorrect,
	}, nil
}

func (m *MeServer) FindAdminUUID(ctx context.Context, req *message.FindAdminUUIDRequest) (*message.FindAdminUUIDResponse, error) {
	tx, err := m.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	defer tx.Rollback()

	fmt.Println("f3 " + req.Uuid)
	isFind, err := m.db.FindAdminUUID(ctx, tx, req.Uuid)
	if err != nil {
		return nil, err
	}

	tx.Commit()
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

	tx.Commit()
	return &message.InsertAdminUUIDResponse{}, nil
}
