package repository

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type admin struct {
}

func (a *admin) GetPassword(ctx context.Context, tx *sqlx.Tx) (string, error) {

	var password []string
	err := tx.Select(&password, "SELECT password FROM me.admin WHERE id = 'admin'")

	if err != nil {
		return "", err
	}

	if len(password) != 1 {
		return "", fmt.Errorf("mysql error. please check the health of the mysql")
	}

	return password[0], nil
}
