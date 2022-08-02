package test

import (
	"github.com/dtc03012/me/db"
	"github.com/dtc03012/me/db/service"
	"github.com/dtc03012/me/protobuf/proto/entity/post"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestDBService_UploadPost(t *testing.T) {
	t.Parallel()

	ctx, tx, _, err := db.SetupMock()
	assert.NoError(t, err)

	postData := &post.Data{
		Title:      "title1",
		Content:    "content1",
		Tags:       []string{"tag1"},
		TimeToRead: 1,
	}

	svc, m := service.NewMockDBService()
	m.PostRepo.On("InsertPost", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil).Once()

	err = svc.UploadPost(ctx, tx, postData)
	assert.NoError(t, err)

	err = svc.UploadPost(ctx, tx, nil)
	assert.Error(t, err)

	m.PostRepo.AssertNumberOfCalls(t, "InsertPost", 1)
	m.PostRepo.AssertExpectations(t)
}
