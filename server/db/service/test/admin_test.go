package test

import (
	"github.com/dtc03012/me/db"
	"github.com/dtc03012/me/db/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestDBService_CheckAdminPassword(t *testing.T) {
	t.Parallel()

	ctx, tx, _, err := db.SetupMock()
	assert.NoError(t, err)

	svc, m := service.NewMockDBService()
	m.AdminRepo.On("GetPassword", mock.Anything, mock.Anything).Return("password", nil).Twice()

	isCorrect, err := svc.CheckAdminPassword(ctx, tx, "password")
	assert.NoError(t, err)
	assert.Equal(t, true, isCorrect)

	isCorrect, err = svc.CheckAdminPassword(ctx, tx, "wrong")
	assert.NoError(t, err)
	assert.Equal(t, false, isCorrect)

	m.AdminRepo.AssertExpectations(t)
}

func TestDBService_InsertAdminUUID(t *testing.T) {
	t.Parallel()

	ctx, tx, _, err := db.SetupMock()
	assert.NoError(t, err)

	svc, m := service.NewMockDBService()
	m.AdminRepo.On("InsertUUID", mock.Anything, mock.Anything, mock.Anything).Return(nil).Once()

	err = svc.InsertAdminUUID(ctx, tx, "password")
	assert.NoError(t, err)

	m.AdminRepo.AssertExpectations(t)
}

func TestDBService_FindAdminUUID(t *testing.T) {
	t.Parallel()

	ctx, tx, _, err := db.SetupMock()
	assert.NoError(t, err)

	svc, m := service.NewMockDBService()
	m.AdminRepo.On("FindUUID", mock.Anything, mock.Anything, mock.Anything).Return("uuid", nil).Once()

	isSame, err := svc.FindAdminUUID(ctx, tx, "uuid")
	assert.NoError(t, err)
	assert.Equal(t, true, isSame)

	m.AdminRepo.AssertExpectations(t)
}

