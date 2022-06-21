package test

import (
	"github.com/dtc03012/me/db"
	"github.com/dtc03012/me/db/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestDbService_CheckAdmin(t *testing.T) {
	t.Parallel()

	ctx, tx, _, err := db.SetupMock()
	assert.NoError(t, err)

	svc, m := service.NewMockDBService()
	m.AdminRepo.On("GetPassword", mock.Anything, mock.Anything).Return("password", nil).Twice()

	isCorrect, err := svc.CheckAdmin(ctx, tx, "password")
	assert.NoError(t, err)
	assert.Equal(t, true, isCorrect)

	isCorrect, err = svc.CheckAdmin(ctx, tx, "wrong")
	assert.NoError(t, err)
	assert.Equal(t, false, isCorrect)

	m.AdminRepo.AssertExpectations(t)
}
