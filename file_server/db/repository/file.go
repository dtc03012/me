package repository

import (
	"context"
	"errors"
	"github.com/jmoiron/sqlx"
)

type file struct {
}

func (f *file) GetFileId(ctx context.Context, tx *sqlx.Tx, fileName string) (string, error) {

	var (
		fileIdList []string
		err        error
	)

	err = tx.SelectContext(ctx, &fileIdList, "SELECT id FROM file WHERE name = ?", fileName)
	if err != nil {
		return "", err
	}

	if len(fileIdList) == 0 {
		return "", errors.New("get file id db repository error: There is no file name in db")
	}

	if len(fileIdList) > 1 {
		return "", errors.New("get file id db repository error: duplicate file name in db")
	}

	return fileIdList[0], nil
}

func (f *file) InsertFileIdName(ctx context.Context, tx *sqlx.Tx, fileName string, fileId string) error {

	_, err := tx.ExecContext(ctx, "INSERT INTO file(id, name) VALUES (?, ?)", fileId, fileName)

	return err
}
