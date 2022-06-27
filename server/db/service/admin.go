package service

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
)

func (dbs *dbService) CheckAdminPassword(ctx context.Context, tx *sqlx.Tx, password string) (bool, error) {
	adminPassword, err := dbs.AdminRepo.GetPassword(ctx, tx)
	if err != nil {
		return false, err
	}

	return adminPassword == password, nil
}

func (dbs *dbService) InsertAdminUUID(ctx context.Context, tx *sqlx.Tx, uuid string) error {
	err := dbs.AdminRepo.InsertUUID(ctx, tx, uuid)
	return err
}

func (dbs *dbService) FindAdminUUID(ctx context.Context, tx *sqlx.Tx, uuid string) (bool, error) {
	expUUID, err := dbs.AdminRepo.FindUUID(ctx, tx, uuid)
	fmt.Println("f1 " + uuid)
	fmt.Println("f2 " + expUUID)
	if err != nil {
		return false, err
	}

	return expUUID == uuid, err
}
