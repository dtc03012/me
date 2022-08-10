package test

import (
	"github.com/dtc03012/me/db"
	"github.com/dtc03012/me/db/entity"
	"github.com/dtc03012/me/db/option"
	"github.com/dtc03012/me/db/service"
	"github.com/dtc03012/me/protobuf/proto/entity/post"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
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
		Views:            1,
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

func TestDBService_FetchPostList(t *testing.T) {
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

	fetchPosts, err := svc.FetchPostList(ctx, tx, 1, 2)
	assert.NoError(t, err)
	assert.Len(t, fetchPosts, 2)
	assert.Equal(t, "title1", fetchPosts[0].Title)
	assert.Equal(t, "title2", fetchPosts[1].Title)

	_, err = svc.FetchPostList(ctx, tx, 0, 2)
	assert.Error(t, err)

	m.PostRepo.AssertNumberOfCalls(t, "GetBulkPost", 1)
	m.PostRepo.AssertExpectations(t)
}

func TestDBService_FetchPost(t *testing.T) {
	t.Parallel()

	ctx, tx, _, err := db.SetupMock()
	assert.NoError(t, err)

	postData := &entity.Post{
		Title:            "title1",
		Content:          "content1",
		Tags:             []string{"tag1"},
		LikeCnt:          3,
		Views:            1,
		TimeToReadMinute: 1,
	}

	svc, m := service.NewMockDBService()
	m.PostRepo.On("GetPost", mock.Anything, mock.Anything, mock.Anything).Return(postData, nil).Once()

	fetchPost, err := svc.FetchPost(ctx, tx, 1)
	assert.NoError(t, err)
	assert.NotNil(t, fetchPost)
	assert.Equal(t, "title1", fetchPost.Title)
	assert.Equal(t, "content1", fetchPost.Content)
	assert.Equal(t, []string{"tag1"}, fetchPost.Tags)
	assert.Equal(t, int32(1), fetchPost.Views)
	assert.Equal(t, int32(3), fetchPost.LikeCnt)
	assert.Equal(t, int32(1), fetchPost.TimeToReadMinute)

	_, err = svc.FetchPost(ctx, tx, 0)
	assert.Error(t, err)

	m.PostRepo.AssertNumberOfCalls(t, "GetPost", 1)
	m.PostRepo.AssertExpectations(t)
}

func TestDBService_IncrementViews(t *testing.T) {
	t.Parallel()

	ctx, tx, _, err := db.SetupMock()
	assert.NoError(t, err)

	svc, m := service.NewMockDBService()
	m.PostRepo.On("InsertViews", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil).Once()

	err = svc.IncrementViews(ctx, tx, 1, "uuid")
	assert.NoError(t, err)

	m.PostRepo.AssertExpectations(t)
}

func TestDBService_LeaveComment(t *testing.T) {
	t.Parallel()

	ctx, tx, _, err := db.SetupMock()
	assert.NoError(t, err)

	comment := &post.Comment{
		Id:       1,
		PostId:   1,
		Writer:   "writer1",
		Password: "password1",
		Comment:  "comment1",
		LikeCnt:  1,
	}

	svc, m := service.NewMockDBService()
	m.PostRepo.On("InsertComment", mock.Anything, mock.Anything, mock.Anything).Return(nil).Once()

	err = svc.LeaveComment(ctx, tx, comment)
	assert.NoError(t, err)

	m.PostRepo.AssertExpectations(t)
}

func TestDBService_FetchCommentList(t *testing.T) {
	t.Parallel()

	ctx, tx, _, err := db.SetupMock()
	assert.NoError(t, err)

	currentTime := time.Now()
	svc, m := service.NewMockDBService()
	m.PostRepo.On("GetBulkComment", mock.Anything, mock.Anything, mock.Anything).Return([]*entity.Comment{
		{
			Id:       1,
			PostId:   1,
			Writer:   "writer1",
			Password: "password1",
			Comment:  "comment1",
			LikeCnt:  1,
			CreateAt: currentTime,
		},
		{
			Id:       2,
			PostId:   1,
			Writer:   "writer2",
			Password: "password2",
			Comment:  "comment2",
			LikeCnt:  1,
			CreateAt: currentTime,
		},
	}, nil)

	commentList, err := svc.FetchCommentList(ctx, tx, &option.CommentOption{
		SizeRange: &option.RangeOption{
			Size: 1,
			Row:  1,
		},
		PostId: 1,
	})
	assert.NoError(t, err)
	assert.NotNil(t, commentList)
	assert.Len(t, commentList, 2)
	assert.Equal(t, "writer1", commentList[0].Writer)
	assert.Equal(t, "writer2", commentList[1].Writer)

	m.PostRepo.AssertExpectations(t)
}

func TestDBService_DeleteComment(t *testing.T) {
	t.Parallel()

	ctx, tx, _, err := db.SetupMock()
	assert.NoError(t, err)

	svc, m := service.NewMockDBService()
	m.PostRepo.On("DeleteComment", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil).Once()

	err = svc.DeleteComment(ctx, tx, 1, 1)
	assert.NoError(t, err)

	err = svc.DeleteComment(ctx, tx, 0, 1)
	assert.Error(t, err)

	err = svc.DeleteComment(ctx, tx, 1, 0)
	assert.Error(t, err)

	m.PostRepo.AssertExpectations(t)
}

func TestDBService_QueryPostList(t *testing.T) {

	t.Parallel()

	ctx, tx, _, err := db.SetupMock()
	assert.NoError(t, err)

	svc, m := service.NewMockDBService()
	m.PostRepo.On("QueryBulkPost", mock.Anything, mock.Anything, mock.Anything).Return([]*entity.Post{
		{
			Id:               1,
			Title:            "title1",
			Content:          "content1",
			LikeCnt:          1,
			TimeToReadMinute: 1,
			Writer:           "writer1",
			Tags:             []string{"tag1"},
		},
	}, nil)

	postList, err := svc.QueryPostList(ctx, tx, &option.PostOption{
		SizeRange: &option.RangeOption{
			Size: 1,
			Row:  1,
		},
		QueryType: option.Title,
		Query:     "title1",
		Tags:      []string{},
	})

	assert.NoError(t, err)
	assert.NotNil(t, postList)
	assert.Len(t, postList, 1)
	assert.Len(t, postList[0].Tags, 1)

	assert.Equal(t, int32(1), postList[0].Id)
	assert.Equal(t, "title1", postList[0].Title)
	assert.Equal(t, "writer1", postList[0].Writer)
	assert.Equal(t, "content1", postList[0].Content)

	m.PostRepo.AssertExpectations(t)
}
