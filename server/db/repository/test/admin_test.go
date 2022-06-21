package test

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/dtc03012/me/db"
	"github.com/dtc03012/me/db/repository"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAdmin_GetPassword(t *testing.T) {
	t.Parallel()

	ctx, tx, mock, err := db.SetupMock()
	assert.NoError(t, err)

	mock.ExpectQuery("SELECT password FROM me.admin WHERE id = 'admin'").
		WillReturnRows(sqlmock.NewRows([]string{"password"}).AddRow("password"))

	admin := repository.NewAdminRepo()
	password, err := admin.GetPassword(ctx, tx)
	assert.Equal(t, "password", password)

	mock.ExpectQuery("SELECT password FROM me.admin WHERE id = 'admin'").
		WillReturnRows(sqlmock.NewRows([]string{"password"}).
			AddRow("password1").AddRow("password2"))

	_, err = admin.GetPassword(ctx, tx)
	assert.Error(t, err)

	mock.ExpectationsWereMet()
}
