package test

import (
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/mysql"
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
	query := goqu.Dialect("mysql").Select("bp.pid", "bp.writer", "bp.title", "bp.content", goqu.COUNT(goqu.I("bl.uuid").Distinct()).As("likes"), "bp.time_to_read_minute", "bp.create_at", goqu.COUNT(goqu.I("bv.uuid").Distinct()).As("views")).
		From(goqu.T("board_post").As("bp")).
		LeftOuterJoin(goqu.T("board_views").As("bv"), goqu.On(goqu.I("bp.pid").Eq(goqu.I("bv.pid")))).
		LeftOuterJoin(goqu.T("board_likes").As("bl"), goqu.On(goqu.I("bp.pid").Eq(goqu.I("bl.pid")))).
		Where(goqu.Ex{"bp.pid": 1}).
		GroupBy("bp.pid")

	expectedSQL, _, err := query.ToSQL()

	mock.ExpectQuery(regexp.QuoteMeta(expectedSQL)).WithArgs().
		WillReturnRows(sqlmock.NewRows([]string{"pid", "writer", "title", "content", "likes", "time_to_read_minute", "create_at", "views"}).AddRow(1, "writer1", "title1", "content1", 3, 1, currentTime, 1))

	postRepo := repository.NewPostRepo()
	post, err := postRepo.GetPost(ctx, tx, 1)
	assert.NoError(t, err)
	assert.NotNil(t, post)

	assert.Equal(t, int32(1), post.Id)
	assert.Equal(t, "writer1", post.Writer)
	assert.Equal(t, "title1", post.Title)
	assert.Equal(t, "content1", post.Content)
	assert.Equal(t, int32(3), post.Likes)
	assert.Equal(t, int32(1), post.TimeToReadMinute)
	assert.Equal(t, currentTime, post.CreateAt)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestPost_GetBulkPost(t *testing.T) {
	t.Parallel()

	var testCase = []struct {
		description string
		opt         *option.PostOption
		postList    []*entity.Post
	}{
		{
			description: "check title and Content",
			opt: &option.PostOption{
				SizeRange: &option.RangeOption{
					Row:  1,
					Size: 1,
				},
				QueryType: option.QueryTitleOrContent,
				Query:     "tc1",
			},
		},
		{
			description: "check title",
			opt: &option.PostOption{
				SizeRange: &option.RangeOption{
					Row:  1,
					Size: 1,
				},
				QueryType: option.QueryTitle,
				Query:     "t1",
			},
		},
		{
			description: "check content",
			opt: &option.PostOption{
				SizeRange: &option.RangeOption{
					Row:  1,
					Size: 1,
				},
				QueryType: option.QueryContent,
				Query:     "c1",
			},
		},
		{
			description: "check writer",
			opt: &option.PostOption{
				SizeRange: &option.RangeOption{
					Row:  1,
					Size: 1,
				},
				QueryType: option.QueryWriter,
				Query:     "w1",
			},
		},
		{
			description: "classification and title check",
			opt: &option.PostOption{
				SizeRange: &option.RangeOption{
					Row:  1,
					Size: 1,
				},
				QueryType:          option.QueryTitle,
				Query:              "t1",
				ClassificationType: option.ClassificationPopular,
			},
		},
	}

	for _, tc := range testCase {
		t.Run(tc.description, func(t *testing.T) {
			ctx, tx, mock, err := db.SetupMock()
			assert.NoError(t, err)

			currentTime := time.Now()

			n, m, err := option.CalculateDBRange(tc.opt.SizeRange)
			assert.NoError(t, err)

			query := goqu.Dialect("mysql").Select("bp.pid", "bp.writer", "bp.title", "bp.content", goqu.COUNT(goqu.I("bl.uuid").Distinct()).As("likes"), "bp.time_to_read_minute", "bp.create_at", goqu.COUNT(goqu.I("bv.uuid").Distinct()).As("views")).
				From(goqu.T("board_post").As("bp")).
				LeftOuterJoin(goqu.T("board_views").As("bv"), goqu.On(goqu.I("bp.pid").Eq(goqu.I("bv.pid")))).
				LeftOuterJoin(goqu.T("board_likes").As("bl"), goqu.On(goqu.I("bp.pid").Eq(goqu.I("bl.pid"))))

			qs := fmt.Sprintf("%%%s%%", tc.opt.Query)

			if tc.opt.QueryType == option.QueryTitleOrContent {
				query = query.Where(
					goqu.ExOr{
						"bp.title":   goqu.Op{"like": qs},
						"bp.content": goqu.Op{"like": qs},
					},
				)
			} else if tc.opt.QueryType == option.QueryTitle {
				query = query.Where(
					goqu.Ex{
						"bp.title": goqu.Op{"like": qs},
					},
				)
			} else if tc.opt.QueryType == option.QueryContent {
				query = query.Where(
					goqu.Ex{
						"bp.content": goqu.Op{"like": qs},
					},
				)
			} else if tc.opt.QueryType == option.QueryWriter {
				query = query.Where(
					goqu.Ex{
						"bp.writer": goqu.Op{"like": qs},
					},
				)
			}

			if tc.opt.ClassificationType == option.ClassificationNotice {
				query = query.Where(goqu.I("bp.is_notice").Eq(true))
			} else {
				query = query.Where(goqu.I("bp.is_notice").Eq(false))
			}

			query = query.GroupBy("bp.pid")

			if tc.opt.ClassificationType == option.ClassificationALL || tc.opt.ClassificationType == option.ClassificationNotice {
				query = query.Order(goqu.I("bp.pid").Desc())
			} else if tc.opt.ClassificationType == option.ClassificationPopular {
				query = query.Order(goqu.I("likes").Desc(), goqu.I("views").Desc(), goqu.I("bp.pid").Desc())
			}

			query = query.Limit(uint(m - n + 1)).Offset(uint(n))

			expectedSQL, _, err := query.ToSQL()
			assert.NoError(t, err)

			mock.ExpectQuery(regexp.QuoteMeta(expectedSQL)).WithArgs().
				WillReturnRows(sqlmock.NewRows([]string{"pid", "writer", "title", "content", "time_to_read_minute", "create_at", "views"}).AddRow(
					1, "writer1", "title1", "content1", 1, currentTime, 1))

			postRepo := repository.NewPostRepo()
			postList, err := postRepo.GetBulkPost(ctx, tx, tc.opt)

			assert.NoError(t, err)
			assert.NotNil(t, postList)
			assert.Len(t, postList, 1)

			assert.Equal(t, int32(1), postList[0].Id)
			assert.Equal(t, "writer1", postList[0].Writer)
			assert.Equal(t, "title1", postList[0].Title)
			assert.Equal(t, "content1", postList[0].Content)
			assert.Equal(t, int32(1), postList[0].TimeToReadMinute)
			assert.Equal(t, currentTime, postList[0].CreateAt)
			assert.Equal(t, int32(1), postList[0].Views)

			err = mock.ExpectationsWereMet()
			assert.NoError(t, err)
		})
	}
}

func TestPost_InsertPost(t *testing.T) {
	t.Parallel()

	ctx, tx, mock, err := db.SetupMock()
	assert.NoError(t, err)

	post := &entity.Post{
		Writer:           "writer1",
		Password:         "password1",
		Title:            "title1",
		Content:          "content1",
		Likes:            3,
		IsNotice:         false,
		TimeToReadMinute: 1,
	}

	expectedSQL := fmt.Sprintf("INSERT IGNORE INTO board_post(writer, password, title, content, is_notice, time_to_read_minute) VALUES (?, ?, ?, ?, ?, ?)")
	mock.ExpectExec(regexp.QuoteMeta(expectedSQL)).WithArgs(post.Writer, post.Password, post.Title, post.Content, post.IsNotice, post.TimeToReadMinute).WillReturnResult(sqlmock.NewResult(1, 1))

	postRepo := repository.NewPostRepo()
	_, err = postRepo.InsertPost(ctx, tx, post)
	assert.NoError(t, err)

	_, err = postRepo.InsertPost(ctx, tx, nil)
	assert.Error(t, err)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestPost_GetViews(t *testing.T) {
	t.Parallel()

	ctx, tx, mock, err := db.SetupMock()
	assert.NoError(t, err)

	expectedSQL := fmt.Sprintf("SELECT count(*) FROM board_views WHERE pid = ?")
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

	expectedSQL := fmt.Sprintf("INSERT IGNORE INTO board_views(pid, uuid) VALUES(?, ?)")
	mock.ExpectExec(regexp.QuoteMeta(expectedSQL)).WithArgs(1, "uuid").WillReturnResult(sqlmock.NewResult(1, 1))

	postRepo := repository.NewPostRepo()
	err = postRepo.InsertViews(ctx, tx, 1, "uuid")
	assert.NoError(t, err)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestPost_GetBulkComment(t *testing.T) {
	t.Parallel()

	ctx, tx, mock, err := db.SetupMock()
	assert.NoError(t, err)

	currentTime := time.Now()
	expectedSQL := fmt.Sprintf("SELECT * FROM board_comment WHERE pid = ? ORDER BY parent_cid DESC, cid ASC LIMIT ?, ?")
	mock.ExpectQuery(regexp.QuoteMeta(expectedSQL)).WithArgs(1, 0, 2).
		WillReturnRows(sqlmock.NewRows([]string{"cid", "pid", "parent_cid", "is_exist", "writer", "password", "comment", "like_cnt", "create_at"}).AddRow(
			1, 1, 1, true, "writer1", "password1", "comment1", 5, currentTime).AddRow(
			2, 1, 1, true, "writer2", "password2", "comment2", 5, currentTime))

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
		Id:              1,
		PostId:          1,
		ParentCommentId: 1,
		IsExist:         true,
		Writer:          "writer1",
		Password:        "password1",
		Comment:         "comment1",
		LikeCnt:         3,
	}

	expectedSQL := fmt.Sprintf("INSERT IGNORE INTO board_comment(pid, writer, parent_cid, is_exist, password, comment, like_cnt) VALUES(?, ?, ?, ?, ?, ?, ?)")
	mock.ExpectExec(regexp.QuoteMeta(expectedSQL)).WithArgs(comment.PostId, comment.Writer, comment.ParentCommentId, comment.IsExist, comment.Password, comment.Comment, comment.LikeCnt).WillReturnResult(sqlmock.NewResult(1, 1))

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

	expectedSQL := fmt.Sprintf("UPDATE board_comment SET is_exist = false WHERE pid = ? and cid = ?")
	mock.ExpectExec(regexp.QuoteMeta(expectedSQL)).WithArgs(1, 1).WillReturnResult(sqlmock.NewResult(1, 1))

	postRepo := repository.NewPostRepo()
	err = postRepo.DeleteComment(ctx, tx, 1, 1)
	assert.NoError(t, err)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestPost_UpdatePost(t *testing.T) {
	t.Parallel()

	ctx, tx, _, err := db.SetupMock()
	assert.NoError(t, err)

	p := &entity.Post{
		Id:               1,
		Title:            "t1",
		Content:          "w1",
		Writer:           "w1",
		TimeToReadMinute: 1,
	}

	postRepo := repository.NewPostRepo()
	err = postRepo.UpdatePost(ctx, tx, p)

}
