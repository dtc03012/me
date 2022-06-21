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

	isCorrect, err := svc.CheckAdmin(ctx, tx, os.Getenv("ME_PASSWORD"))
	assert.NoError(t, err)
	assert.Equal(t, true, isCorrect)

	isCorrect, err = svc.CheckAdmin(ctx, tx, "wrong")
	assert.NoError(t, err)
	assert.Equal(t, false, isCorrect)
}
