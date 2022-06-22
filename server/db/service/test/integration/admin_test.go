package integration

import (
	"context"
	"github.com/dtc03012/me/db/service"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestIntegration_CheckAdmin(t *testing.T) {
	ctx := context.Background()
	svc := service.NewDBService()
	tx, err := svc.BeginTx(ctx, nil)
	assert.NoError(t, err)

	isCorrect, err := svc.CheckAdminPassword(ctx, tx, os.Getenv("ME_ADMIN_PASSWORD"))
	assert.NoError(t, err)
	assert.Equal(t, true, isCorrect)

	isCorrect, err = svc.CheckAdminPassword(ctx, tx, "wrong")
	assert.NoError(t, err)
	assert.Equal(t, false, isCorrect)
}

func TestIntegration_AdminUUID(t *testing.T) {
	ctx := context.Background()
	svc := service.NewDBService()
	tx, err := svc.BeginTx(ctx, nil)
	assert.NoError(t, err)

	err = svc.InsertAdminUUID(ctx, tx, "uuid")
	assert.NoError(t, err)

	isSame, err := svc.FindAdminUUID(ctx, tx, "uuid")
	assert.NoError(t, err)
	assert.Equal(t, true, isSame)

	isSame, err = svc.FindAdminUUID(ctx, tx, "wrong")
	assert.NoError(t, err)
	assert.Equal(t, false, isSame)
}
