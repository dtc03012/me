package service

import (
	"github.com/dtc03012/me/db"
	"github.com/dtc03012/me/db/repository/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestDbService_CheckAdmin(t *testing.T) {
	t.Parallel()

	ctx, tx, _, err := db.SetupMock()
	assert.NoError(t, err)

	svc := NewMockDBService()
	m := &mocks.Admin{}
	m.On("GetPassword", mock.Anything, mock.Anything).Return("password", nil).Twice()
	svc.AdminRepo = m

	isCorrect, err := svc.CheckAdmin(ctx, tx, "password")
	assert.NoError(t, err)
	assert.Equal(t, true, isCorrect)

	isCorrect, err = svc.CheckAdmin(ctx, tx, "wrong")
	assert.NoError(t, err)
	assert.Equal(t, false, isCorrect)

	m.AssertExpectations(t)
}
