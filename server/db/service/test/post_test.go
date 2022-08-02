package test

import (
	"github.com/dtc03012/me/db"
	"github.com/dtc03012/me/db/entity"
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
		Title:            "title1",
		Content:          "content1",
		Tags:             []string{"tag1"},
		LikeCnt:          3,
		TimeToReadMinute: 1,
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

func TestDBService_FetchPosts(t *testing.T) {
	t.Parallel()

	ctx, tx, _, err := db.SetupMock()
	assert.NoError(t, err)

	postData1 := &entity.Post{
		Title:            "title1",
		Content:          "content1",
		Tags:             []string{"tag1"},
		LikeCnt:          3,
		TimeToReadMinute: 1,
	}

	postData2 := &entity.Post{
		Title:            "title2",
		Content:          "content2",
		Tags:             []string{"tag2"},
		LikeCnt:          3,
		TimeToReadMinute: 1,
	}

	posts := make([]*entity.Post, 0, 1)
	posts = append(posts, postData1)
	posts = append(posts, postData2)

	svc, m := service.NewMockDBService()
	m.PostRepo.On("GetBulkPost", mock.Anything, mock.Anything, mock.Anything).Return(posts, nil).Once()

	fetchPosts, err := svc.FetchPosts(ctx, tx, 1, 2)
	assert.NoError(t, err)
	assert.Len(t, fetchPosts, 2)
	assert.Equal(t, "title1", fetchPosts[0].Title)
	assert.Equal(t, "title2", fetchPosts[1].Title)

	_, err = svc.FetchPosts(ctx, tx, 0, 2)
	assert.Error(t, err)

	m.PostRepo.AssertNumberOfCalls(t, "GetBulkPost", 1)
	m.PostRepo.AssertExpectations(t)
}
