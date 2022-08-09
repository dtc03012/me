package service

import (
	"context"
	"errors"
	"github.com/jmoiron/sqlx"
)

func (dbs *dbService) GetFileId(ctx context.Context, tx *sqlx.Tx, fileName string) (string, error) {

	if len(fileName) == 0 {
		return "", errors.New("get file id db service error: file name is empty")
	}

	fileId, err := dbs.FileRepo.GetFileId(ctx, tx, fileName)
	if err != nil {
		return "", err
	}

	return fileId, nil
}

func (dbs *dbService) InsertFileInfo(ctx context.Context, tx *sqlx.Tx, fileName string, fileId string) error {

	if len(fileName) == 0 || len(fileId) == 0 {
		return errors.New("get file id db service error: file name or id is empty")
	}

	err := dbs.FileRepo.InsertFileIdName(ctx, tx, fileName, fileId)

	return err
}
