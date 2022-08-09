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

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestPost_GetBulkPost(t *testing.T) {
	t.Parallel()

	ctx, tx, mock, err := db.SetupMock()
	assert.NoError(t, err)

	currentTime := time.Now()
	expectedSQL := fmt.Sprintf("SELECT * FROM board_post ORDER BY pid DESC LIMIT ?, ?")
	mock.ExpectQuery(regexp.QuoteMeta(expectedSQL)).WithArgs(0, 1).
		WillReturnRows(sqlmock.NewRows([]string{"pid", "writer", "title", "content", "like_cnt", "views", "time_to_read_minute", "create_at"}).AddRow(1, "writer1", "title1", "content1", 3, 1, 1, currentTime))

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
	assert.Equal(t, int32(1), post[0].Views)
	assert.Equal(t, int32(1), post[0].TimeToReadMinute)
	assert.Equal(t, currentTime, post[0].CreateAt)

	assert.Len(t, post[0].Tags, 1)
	assert.Equal(t, "tag1", post[0].Tags[0])

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestPost_InsertPost(t *testing.T) {
	t.Parallel()

	ctx, tx, mock, err := db.SetupMock()
	assert.NoError(t, err)

	post := &entity.Post{
		Writer:           "writer1",
		Title:            "title1",
		Content:          "content1",
		LikeCnt:          3,
		Views:            1,
		TimeToReadMinute: 1,
	}

	tags := []string{"tag1"}

	expectedSQL := fmt.Sprintf("INSERT IGNORE INTO board_post(writer, title, content, like_cnt, views, time_to_read_minute) VALUES (?, ?, ?, ?, ?, ?)")
	mock.ExpectExec(regexp.QuoteMeta(expectedSQL)).WithArgs(post.Writer, post.Title, post.Content, post.LikeCnt, post.Views, post.TimeToReadMinute).WillReturnResult(sqlmock.NewResult(1, 1))

	expectedSQL = fmt.Sprintf("INSERT IGNORE INTO board_tag(value) VALUES (?)")
	mock.ExpectExec(regexp.QuoteMeta(expectedSQL)).WithArgs(tags[0]).WillReturnResult(sqlmock.NewResult(1, 1))

	expectedSQL = fmt.Sprintf("INSERT IGNORE INTO board_post_tag(tid, pid) VALUES(?, ?)")
	mock.ExpectExec(regexp.QuoteMeta(expectedSQL)).WithArgs(1, 1).WillReturnResult(sqlmock.NewResult(1, 1))

	postRepo := repository.NewPostRepo()
	err = postRepo.InsertPost(ctx, tx, post, tags)
	assert.NoError(t, err)

	err = postRepo.InsertPost(ctx, tx, nil, tags)
	assert.Error(t, err)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestPost_GetViews(t *testing.T) {
	t.Parallel()

	ctx, tx, mock, err := db.SetupMock()
	assert.NoError(t, err)

	expectedSQL := fmt.Sprintf("SELECT views FROM board_post WHERE pid = ?")
	mock.ExpectQuery(regexp.QuoteMeta(expectedSQL)).WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"views"}).AddRow(1))

	postRepo := repository.NewPostRepo()
	post, err := postRepo.GetViews(ctx, tx, 1)
	assert.NoError(t, err)
	assert.Equal(t, 1, post)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestPost_UpdateViews(t *testing.T) {
	t.Parallel()

	ctx, tx, mock, err := db.SetupMock()
	assert.NoError(t, err)

	expectedSQL := fmt.Sprintf("UPDATE board_post SET views = ? WHERE pid = ?")
	mock.ExpectExec(regexp.QuoteMeta(expectedSQL)).WithArgs(1, 1).WillReturnResult(sqlmock.NewResult(1, 1))

	postRepo := repository.NewPostRepo()
	err = postRepo.UpdateViews(ctx, tx, 1, 1)
	assert.NoError(t, err)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestPost_GetBulkComment(t *testing.T) {
	t.Parallel()

	ctx, tx, mock, err := db.SetupMock()
	assert.NoError(t, err)

	currentTime := time.Now()
	expectedSQL := fmt.Sprintf("SELECT * FROM board_comment WHERE pid = ? ORDER BY cid LIMIT ?, ?")
	mock.ExpectQuery(regexp.QuoteMeta(expectedSQL)).WithArgs(1, 0, 2).
		WillReturnRows(sqlmock.NewRows([]string{"cid", "pid", "writer", "password", "comment", "like_cnt", "create_at"}).AddRow(
			1, 1, "writer1", "password1", "comment1", 5, currentTime).AddRow(
			2, 1, "writer2", "password2", "comment2", 5, currentTime))

	postRepo := repository.NewPostRepo()
	comments, err := postRepo.GetBulkComment(ctx, tx, &option.CommentOption{
		SizeRange: &option.RangeOption{
			Row:  1,
			Size: 2,
		},
		PostId: 1,
	})

	assert.NoError(t, err)
	assert.NotNil(t, comments)
	assert.Len(t, comments, 2)

	assert.Equal(t, int32(1), comments[0].Id)
	assert.Equal(t, int32(2), comments[1].Id)
	assert.Equal(t, int32(1), comments[0].PostId)
	assert.Equal(t, int32(1), comments[1].PostId)
	assert.Equal(t, "comment1", comments[0].Comment)
	assert.Equal(t, "comment2", comments[1].Comment)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestPost_InsertComment(t *testing.T) {
	t.Parallel()

	ctx, tx, mock, err := db.SetupMock()
	assert.NoError(t, err)

	comment := &entity.Comment{
		Id:       1,
		PostId:   1,
		Writer:   "writer1",
		Password: "password1",
		Comment:  "comment1",
		LikeCnt:  3,
	}

	expectedSQL := fmt.Sprintf("INSERT IGNORE INTO board_comment(pid, writer, password, comment, like_cnt) VALUES(?, ?, ?, ?, ?)")
	mock.ExpectExec(regexp.QuoteMeta(expectedSQL)).WithArgs(comment.PostId, comment.Writer, comment.Password, comment.Comment, comment.LikeCnt).WillReturnResult(sqlmock.NewResult(1, 1))

	postRepo := repository.NewPostRepo()
	err = postRepo.InsertComment(ctx, tx, comment)
	assert.NoError(t, err)

	err = postRepo.InsertComment(ctx, tx, nil)
	assert.Error(t, err)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestPost_DeleteComment(t *testing.T) {
	t.Parallel()

	ctx, tx, mock, err := db.SetupMock()
	assert.NoError(t, err)

	expectedSQL := fmt.Sprintf("DELETE FROM board_comment WHERE pid = ? and cid = ?")
	mock.ExpectExec(regexp.QuoteMeta(expectedSQL)).WithArgs(1, 1).WillReturnResult(sqlmock.NewResult(1, 1))

	postRepo := repository.NewPostRepo()
	err = postRepo.DeleteComment(ctx, tx, 1, 1)
	assert.NoError(t, err)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}
