package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/mysql"
	"github.com/dtc03012/me/db/entity"
	"github.com/dtc03012/me/db/option"
	"github.com/jmoiron/sqlx"
)

type post struct {
}

var mysql = goqu.Dialect("mysql")

func (a *post) GetPost(ctx context.Context, tx *sqlx.Tx, pid int32) (*entity.Post, error) {

	var post []*entity.Post
	post = make([]*entity.Post, 0)

	query := mysql.Select("bp.pid", "bp.writer", "bp.title", "bp.content", goqu.COUNT(goqu.I("bl.uuid").Distinct()).As("likes"), "bp.time_to_read_minute", "bp.create_at", goqu.COUNT(goqu.I("bv.uuid").Distinct()).As("views")).
		From(goqu.T("board_post").As("bp")).
		LeftOuterJoin(goqu.T("board_views").As("bv"), goqu.On(goqu.I("bp.pid").Eq(goqu.I("bv.pid")))).
		LeftOuterJoin(goqu.T("board_likes").As("bl"), goqu.On(goqu.I("bp.pid").Eq(goqu.I("bl.pid")))).
		Where(goqu.Ex{"bp.pid": pid}).
		GroupBy("bp.pid")

	sql, _, err := query.ToSQL()

	err = tx.SelectContext(ctx, &post, sql)
	if err != nil {
		return nil, err
	}

	if len(post) == 0 {
		return nil, errors.New("post db repository error: there is no file id in db")
	}

	if len(post) > 1 {
		return nil, errors.New("post db repository error: duplicate file id. it is caused by server error")
	}

	return post[0], nil
}

func (a *post) GetBulkPost(ctx context.Context, tx *sqlx.Tx, opt *option.PostOption) ([]*entity.Post, error) {

	var postList []*entity.Post

	postList = make([]*entity.Post, 0)

	n, m, err := option.CalculateDBRange(opt.SizeRange)
	if err != nil {
		return nil, err
	}

	query := mysql.Select("bp.pid", "bp.writer", "bp.title", "bp.content", goqu.COUNT(goqu.I("bl.uuid").Distinct()).As("likes"), "bp.time_to_read_minute", "bp.create_at", goqu.COUNT(goqu.I("bv.uuid").Distinct()).As("views")).
		From(goqu.T("board_post").As("bp")).
		LeftOuterJoin(goqu.T("board_views").As("bv"), goqu.On(goqu.I("bp.pid").Eq(goqu.I("bv.pid")))).
		LeftOuterJoin(goqu.T("board_likes").As("bl"), goqu.On(goqu.I("bp.pid").Eq(goqu.I("bl.pid"))))

	qs := fmt.Sprintf("%%%s%%", opt.Query)

	if opt.QueryType == option.QueryTitleOrContent {
		query = query.Where(
			goqu.ExOr{
				"bp.title":   goqu.Op{"like": qs},
				"bp.content": goqu.Op{"like": qs},
			},
		)
	} else if opt.QueryType == option.QueryTitle {
		query = query.Where(
			goqu.Ex{
				"bp.title": goqu.Op{"like": qs},
			},
		)
	} else if opt.QueryType == option.QueryContent {
		query = query.Where(
			goqu.Ex{
				"bp.content": goqu.Op{"like": qs},
			},
		)
	} else if opt.QueryType == option.QueryWriter {
		query = query.Where(
			goqu.Ex{
				"bp.writer": goqu.Op{"like": qs},
			},
		)
	}

	if opt.ClassificationType == option.ClassificationNotice {
		query = query.Where(goqu.I("bp.is_notice").Eq(true))
	} else {
		query = query.Where(goqu.I("bp.is_notice").Eq(false))
	}

	query = query.GroupBy("bp.pid")

	if opt.ClassificationType == option.ClassificationALL || opt.ClassificationType == option.ClassificationNotice {
		query = query.Order(goqu.I("bp.pid").Desc())
	} else if opt.ClassificationType == option.ClassificationPopular {
		query = query.Order(goqu.I("likes").Desc(), goqu.I("views").Desc(), goqu.I("bp.pid").Desc())
	}

	query = query.Limit(uint(m - n + 1)).Offset(uint(n))

	sql, _, err := query.ToSQL()
	if err != nil {
		return nil, err
	}

	err = tx.SelectContext(ctx, &postList, sql)

	if err != nil {
		return nil, err
	}

	return postList, nil
}

func (a *post) GetBulkTag(ctx context.Context, tx *sqlx.Tx, pid int32) ([]string, error) {

	var tagList []string

	query := mysql.Select("value").
		From("board_tag").
		Where(goqu.I("board_tag.tid").
			In(
				mysql.Select("board_post_tag.tid").
					From("board_post_tag").
					Where(goqu.Ex{"board_post_tag.pid": pid})))

	sql, _, err := query.ToSQL()
	if err != nil {
		return nil, err
	}

	err = tx.SelectContext(ctx, &tagList, sql)
	if err != nil {
		return nil, err
	}

	return tagList, nil
}

func (a *post) InsertPost(ctx context.Context, tx *sqlx.Tx, post *entity.Post) (int, error) {
	if post == nil {
		return 0, errors.New("post db repository error: post is nil")
	}

	postResult, err := tx.ExecContext(ctx, "INSERT IGNORE INTO board_post(writer, title, content, is_notice, time_to_read_minute) VALUES (?, ?, ?, ?, ?)", post.Writer, post.Title, post.Content, post.IsNotice, post.TimeToReadMinute)
	if err != nil {
		return 0, err
	}

	pid, err := postResult.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(pid), nil
}

func (a *post) DeletePost(ctx context.Context, tx *sqlx.Tx, pid int32) error {

	query := mysql.Delete("board_post").Where(goqu.Ex{
		"pid": pid,
	})

	sql, _, err := query.ToSQL()
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, sql)
	return err
}

func (a *post) UpdatePost(ctx context.Context, tx *sqlx.Tx, post *entity.Post) error {

	query := mysql.Update("board_post").Set(entity.Post{
		Title:   post.Title,
		Content: post.Content,
	}).Where(goqu.Ex{
		"pid": post.Id,
	})

	sql, _, err := query.ToSQL()
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, sql)
	return err
}

func (a *post) GetTotalPostCount(ctx context.Context, tx *sqlx.Tx) (int, error) {

	var totalCount []int

	err := tx.SelectContext(ctx, &totalCount, "SELECT COUNT(*) FROM board_post")
	if err != nil {
		return 0, err
	}

	if len(totalCount) != 1 {
		return 0, errors.New("get total post count db repository error: unexpected error")
	}

	return totalCount[0], nil
}

func (a *post) InsertBulkTag(ctx context.Context, tx *sqlx.Tx, pid int32, tags []string) error {

	for _, tag := range tags {
		tagResult, err := tx.ExecContext(ctx, "INSERT IGNORE INTO board_tag(value) VALUES (?)", tag)
		if err != nil {
			return err
		}

		tid, err := tagResult.LastInsertId()
		if err != nil {
			return err
		}

		_, err = tx.ExecContext(ctx, "INSERT IGNORE INTO board_post_tag(tid, pid) VALUES(?, ?)", tid, pid)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *post) DeleteBulkTag(ctx context.Context, tx *sqlx.Tx, pid int32) error {

	query := mysql.Delete("board_post_tag").Where(goqu.Ex{
		"pid": pid,
	})

	sql, _, err := query.ToSQL()
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, sql)
	return err
}

func (a *post) GetViews(ctx context.Context, tx *sqlx.Tx, pid int32) (int, error) {

	var views []int32

	err := tx.SelectContext(ctx, &views, "SELECT count(*) FROM board_views WHERE pid = ?", pid)
	if err != nil {
		return 0, err
	}

	if len(views) == 0 {
		return 0, errors.New("get views db repository error: there is no given pid in db")
	}

	return int(views[0]), nil
}

func (a *post) InsertViews(ctx context.Context, tx *sqlx.Tx, pid int32, uuid string) error {

	_, err := tx.ExecContext(ctx, "INSERT IGNORE INTO board_views(pid, uuid) VALUES(?, ?)", pid, uuid)
	if err != nil {
		return err
	}

	return nil
}

func (a *post) GetBulkComment(ctx context.Context, tx *sqlx.Tx, opt *option.CommentOption) ([]*entity.Comment, error) {

	var commentList []*entity.Comment
	commentList = make([]*entity.Comment, 0)

	n, m, err := option.CalculateDBRange(opt.SizeRange)
	if err != nil {
		return nil, err
	}

	err = tx.SelectContext(ctx, &commentList, "SELECT * FROM board_comment WHERE pid = ? ORDER BY parent_cid DESC, cid ASC LIMIT ?, ?", opt.PostId, n, m)

	if err != nil {
		return nil, err
	}

	return commentList, nil
}

func (a *post) InsertComment(ctx context.Context, tx *sqlx.Tx, comment *entity.Comment) error {

	if comment == nil {
		return errors.New("insert comment db repository error: comment is nil")
	}

	if comment.ParentCommentId == 0 {
		result, err := tx.ExecContext(ctx, "INSERT IGNORE INTO board_comment(pid, writer, is_exist, password, comment, like_cnt) VALUES(?, ?, ?, ?, ?, ?)",
			comment.PostId, comment.Writer, comment.IsExist, comment.Password, comment.Comment, comment.LikeCnt)
		if err != nil {
			return err
		}

		lastInsertID, err := result.LastInsertId()
		if err != nil {
			return err
		}

		_, err = tx.ExecContext(ctx, "UPDATE board_comment SET parent_cid = ? WHERE cid = ?", lastInsertID, lastInsertID)
		if err != nil {
			return err
		}
	} else {
		_, err := tx.ExecContext(ctx, "INSERT IGNORE INTO board_comment(pid, writer, parent_cid, is_exist, password, comment, like_cnt) VALUES(?, ?, ?, ?, ?, ?, ?)",
			comment.PostId, comment.Writer, comment.ParentCommentId, comment.IsExist, comment.Password, comment.Comment, comment.LikeCnt)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *post) DeleteComment(ctx context.Context, tx *sqlx.Tx, postId int, commentId int) error {

	_, err := tx.ExecContext(ctx, "UPDATE board_comment SET is_exist = false WHERE pid = ? and cid = ?", postId, commentId)
	if err != nil {
		return err
	}

	return nil
}

func (a *post) GetTotalCommentCount(ctx context.Context, tx *sqlx.Tx, pid int32) (int, error) {

	var totalCount []int

	err := tx.SelectContext(ctx, &totalCount, "SELECT COUNT(*) FROM board_comment WHERE pid = ?", pid)
	if err != nil {
		return 0, err
	}

	if len(totalCount) != 1 {
		return 0, errors.New("get total comment count db error: unexpected error")
	}

	return totalCount[0], nil
}

func (a *post) CheckUserLike(ctx context.Context, tx *sqlx.Tx, pid int32, uuid string) (bool, error) {

	var likeList []int

	query := mysql.Select(goqu.COUNT("*")).From("board_likes").Where(goqu.Ex{
		"pid":  pid,
		"uuid": uuid,
	})

	sql, _, err := query.ToSQL()
	if err != nil {
		return false, err
	}

	err = tx.SelectContext(ctx, &likeList, sql)
	if err != nil {
		return false, err
	}

	if len(likeList) != 1 {
		return false, errors.New("check user like db repository error: unexpected error")
	}

	return likeList[0] == 1, nil
}

func (a *post) InsertLike(ctx context.Context, tx *sqlx.Tx, pid int32, uuid string) error {

	_, err := tx.ExecContext(ctx, "INSERT IGNORE INTO board_likes VALUES (?, ?)", pid, uuid)
	return err
}

func (a *post) DeleteLike(ctx context.Context, tx *sqlx.Tx, pid int32, uuid string) error {

	query := mysql.Delete("board_likes").Where(goqu.Ex{
		"pid":  pid,
		"uuid": uuid,
	})

	sql, _, err := query.ToSQL()
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, sql)
	return err
}

func (a *post) CheckPostPassword(ctx context.Context, tx *sqlx.Tx, pid int32, password string) (bool, error) {

	var pidList []int

	query := mysql.Select("pid").From("board_post").Where(goqu.Ex{
		"pid":      pid,
		"password": password,
	})

	sql, _, err := query.ToSQL()
	if err != nil {
		return false, err
	}

	err = tx.SelectContext(ctx, &pidList, sql)
	if err != nil {
		return false, err
	}

	if len(pidList) > 1 {
		return false, errors.New("check post password repository error: unexpected error")
	}

	if len(pidList) == 0 {
		return false, nil
	}

	return true, nil
}
