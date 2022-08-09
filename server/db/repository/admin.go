package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type admin struct {
}

func (a *admin) GetPassword(ctx context.Context, tx *sqlx.Tx) (string, error) {

	var password []string
	err := tx.Select(&password, "SELECT password FROM admin WHERE id = 'admin'")

	if err != nil {
		return "", err
	}

	if len(password) != 1 {
		return "", fmt.Errorf("admin db repository error:  please check the health of the mysql")
	}

	return password[0], nil
}

func (a *admin) InsertUUID(ctx context.Context, tx *sqlx.Tx, uuid string) error {

	if len(uuid) == 0 {
		return errors.New("admin db repository error: uuid isn't set")
	}

	_, err := tx.Exec("INSERT IGNORE INTO admin_login_list VALUES (?)", uuid)

	if err != nil {
		return err
	}

	return nil
}

func (a *admin) FindUUID(ctx context.Context, tx *sqlx.Tx, uuid string) (string, error) {

	if len(uuid) == 0 {
		return "", errors.New("admin db repository error: uuid isn't set")
	}

	var uuidList []string
	err := tx.Select(&uuidList, "SELECT * FROM admin_login_list WHERE uuid = ?", uuid)

	if err != nil {
		return "", err
	}

	if len(uuidList) == 0 {
		return "", nil
	}

	return uuidList[0], nil

}
