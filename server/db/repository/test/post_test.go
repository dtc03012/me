package test

import (
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/dtc03012/me/db"
	"github.com/dtc03012/me/db/entity"
	"github.com/dtc03012/me/db/option"
	"github.com/dtc03012/me/db/repository"
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
	"time"
)

func TestPost_GetPost(t *testing.T) {
	t.Parallel()

	ctx, tx, mock, err := db.SetupMock()
	assert.NoError(t, err)

	currentTime := time.Now()
	expectedSQL := fmt.Sprintf("SELECT * FROM board_post WHERE pid = ?")
	mock.ExpectQuery(regexp.QuoteMeta(expectedSQL)).WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"pid", "writer", "title", "content", "like_cnt", "time_to_read_minute", "create_at"}).AddRow(1, "writer1", "title1", "content1", 3, 1, currentTime))

	expectedSQL = fmt.Sprintf("SELECT value FROM board_tag WHERE board_tag.tid IN (SELECT board_post_tag.tid FROM board_post_tag WHERE board_post_tag.pid = ?)")
	mock.ExpectQuery(regexp.QuoteMeta(expectedSQL)).WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"value"}).AddRow("tag1"))

	postRepo := repository.NewPostRepo()
	post, err := postRepo.GetPost(ctx, tx, 1)
	assert.NoError(t, err)
	assert.NotNil(t, post)
	assert.NotNil(t, post.Tags)

	assert.Equal(t, int32(1), post.Id)
	assert.Equal(t, "writer1", post.Writer)
	assert.Equal(t, "title1", post.Title)
	assert.Equal(t, "content1", post.Content)
	assert.Equal(t, int32(3), post.LikeCnt)
	assert.Equal(t, int32(1), post.TimeToReadMinute)
	assert.Equal(t, currentTime, post.CreateAt)

	assert.Len(t, post.Tags, 1)
	assert.Equal(t, "tag1", post.Tags[0])

	mock.ExpectationsWereMet()
}

func TestPost_GetBulkPost(t *testing.T) {
	t.Parallel()

	ctx, tx, mock, err := db.SetupMock()
	assert.NoError(t, err)

	currentTime := time.Now()
	expectedSQL := fmt.Sprintf("SELECT * FROM board_post ORDER BY pid DESC LIMIT ?, ?")
	mock.ExpectQuery(regexp.QuoteMeta(expectedSQL)).WithArgs(0, 1).
		WillReturnRows(sqlmock.NewRows([]string{"pid", "writer", "title", "content", "like_cnt", "time_to_read_minute", "create_at"}).AddRow(1, "writer1", "title1", "content1", 3, 1, currentTime))

	expectedSQL = fmt.Sprintf("SELECT value FROM board_tag WHERE board_tag.tid IN (SELECT board_post_tag.tid FROM board_post_tag WHERE board_post_tag.pid = ?)")
	mock.ExpectQuery(regexp.QuoteMeta(expectedSQL)).WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"value"}).AddRow("tag1"))

	postRepo := repository.NewPostRepo()
	post, err := postRepo.GetBulkPost(ctx, tx, &option.PostOption{SizeRange: &option.RangeOption{Row: 1, Size: 1}})
	assert.NoError(t, err)
	assert.NotNil(t, post)
	assert.Len(t, post, 1)
	assert.NotNil(t, post[0].Tags)

	assert.Equal(t, int32(1), post[0].Id)
	assert.Equal(t, "writer1", post[0].Writer)
	assert.Equal(t, "title1", post[0].Title)
	assert.Equal(t, "content1", post[0].Content)
	assert.Equal(t, int32(3), post[0].LikeCnt)
	assert.Equal(t, int32(1), post[0].TimeToReadMinute)
	assert.Equal(t, currentTime, post[0].CreateAt)

	assert.Len(t, post[0].Tags, 1)
	assert.Equal(t, "tag1", post[0].Tags[0])

	mock.ExpectationsWereMet()
}

func TestPost_InsertPost(t *testing.T) {
	t.Parallel()

	ctx, tx, mock, err := db.SetupMock()
	assert.NoError(t, err)

	post := &entity.Post{
		Writer:           "writer1",
		Title:            "title1",
		Content:          "content1",
		TimeToReadMinute: 1,
	}

	tags := []string{"tag1"}

	expectedSQL := fmt.Sprintf("INSERT IGNORE INTO board_post(writer, title, content, like_cnt, time_to_read_minute) VALUES (?, ?, ?, ?, ?)")
	mock.ExpectExec(regexp.QuoteMeta(expectedSQL)).WithArgs(post.Writer, post.Title, post.Content, post.LikeCnt, post.TimeToReadMinute).WillReturnResult(sqlmock.NewResult(1, 1))

	expectedSQL = fmt.Sprintf("INSERT IGNORE INTO board_tag(value) VALUES (?)")
	mock.ExpectExec(regexp.QuoteMeta(expectedSQL)).WithArgs(tags[0]).WillReturnResult(sqlmock.NewResult(1, 1))

	expectedSQL = fmt.Sprintf("INSERT IGNORE INTO board_post_tag(tid, pid) VALUES(?, ?)")
	mock.ExpectExec(regexp.QuoteMeta(expectedSQL)).WithArgs(1, 1).WillReturnResult(sqlmock.NewResult(1, 1))

	postRepo := repository.NewPostRepo()
	err = postRepo.InsertPost(ctx, tx, post, tags)
	assert.NoError(t, err)

	err = postRepo.InsertPost(ctx, tx, nil, tags)
	assert.Error(t, err)

	mock.ExpectationsWereMet()
}
