package test

import (
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/dtc03012/me/db"
	"github.com/dtc03012/me/db/repository"
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
)

func TestAdmin_GetPassword(t *testing.T) {
	t.Parallel()

	ctx, tx, mock, err := db.SetupMock()
	assert.NoError(t, err)

	mock.ExpectQuery("SELECT password FROM admin WHERE id = 'admin'").
		WillReturnRows(sqlmock.NewRows([]string{"password"}).AddRow("password"))

	admin := repository.NewAdminRepo()
	password, err := admin.GetPassword(ctx, tx)
	assert.Equal(t, "password", password)

	mock.ExpectQuery("SELECT password FROM admin WHERE id = 'admin'").
		WillReturnRows(sqlmock.NewRows([]string{"password"}).
			AddRow("password1").AddRow("password2"))

	_, err = admin.GetPassword(ctx, tx)
	assert.Error(t, err)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestAdmin_InsertUUID(t *testing.T) {
	t.Parallel()

	ctx, tx, mock, err := db.SetupMock()
	assert.NoError(t, err)

	expectedSQL := fmt.Sprintf("INSERT IGNORE INTO admin_login_list VALUES (?)")
	mock.ExpectExec(regexp.QuoteMeta(expectedSQL)).WithArgs("uuid").WillReturnResult(sqlmock.NewResult(1, 1))

	adminRepo := repository.NewAdminRepo()
	err = adminRepo.InsertUUID(ctx, tx, "uuid")
	assert.NoError(t, err)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestAdmin_FindUUID(t *testing.T) {
	t.Parallel()

	ctx, tx, mock, err := db.SetupMock()
	assert.NoError(t, err)

	expectedSQL := fmt.Sprintf("SELECT * FROM admin_login_list WHERE uuid = ?")
	mock.ExpectQuery(regexp.QuoteMeta(expectedSQL)).WithArgs("uuid").
		WillReturnRows(sqlmock.NewRows([]string{"uuid"}).AddRow("uuid"))

	adminRepo := repository.NewAdminRepo()
	uuid, err := adminRepo.FindUUID(ctx, tx, "uuid")
	assert.NoError(t, err)
	assert.Equal(t, "uuid", uuid)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}
