package service

import (
	"context"
	"github.com/jmoiron/sqlx"
)

func (dbs *dbService) CheckAdmin(ctx context.Context, tx *sqlx.Tx, password string) (bool, error) {
	adminPassword, err := dbs.AdminRepo.GetPassword(ctx, tx)
	if err != nil {
		return false, err
	}

	return adminPassword == password, nil
}
