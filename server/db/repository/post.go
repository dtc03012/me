package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/dtc03012/me/db/entity"
	"github.com/dtc03012/me/db/option"
	"github.com/jmoiron/sqlx"
)

type post struct {
}

func (a *post) GetPost(ctx context.Context, tx *sqlx.Tx, pid int32) (*entity.Post, error) {

	var post []*entity.Post
	post = make([]*entity.Post, 0)

	err := tx.SelectContext(ctx, &post, "SELECT bp.pid, bp.writer, bp.title, bp.content, bp.like_cnt, bp.time_to_read_minute, bp.create_at, COUNT(*) as views FROM board_post as bp LEFT OUTER JOIN board_views as bv ON bp.pid = bv.pid WHERE bp.pid = ? GROUP BY bp.pid", pid)

	if err != nil {
		return nil, err
	}

	if len(post) == 0 {
		return nil, errors.New("post db repository error: there is no file id in db")
	}

	if len(post) > 1 {
		return nil, errors.New("post db repository error: duplicate file id. it is caused by server error")
	}

	err = tx.SelectContext(ctx, &post[0].Tags, "SELECT value FROM board_tag WHERE board_tag.tid IN (SELECT board_post_tag.tid FROM board_post_tag WHERE board_post_tag.pid = ?)", pid)

	if err != nil {
		return nil, err
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

	err = tx.SelectContext(ctx, &postList, "SELECT bp.pid, bp.writer, bp.title, bp.content, bp.like_cnt, bp.time_to_read_minute, bp.create_at, COUNT(*) as views FROM board_post as bp LEFT OUTER JOIN board_views as bv ON bp.pid = bv.pid GROUP BY bp.pid ORDER BY pid DESC LIMIT ?, ?", n, m)

	if err != nil {
		return nil, err
	}

	for _, post := range postList {
		err = tx.SelectContext(ctx, &post.Tags, "SELECT value FROM board_tag WHERE board_tag.tid IN (SELECT board_post_tag.tid FROM board_post_tag WHERE board_post_tag.pid = ?)", post.Id)
	}

	if err != nil {
		return nil, err
	}

	return postList, nil
}

func (a *post) InsertPost(ctx context.Context, tx *sqlx.Tx, post *entity.Post, tags []string) error {
	if post == nil {
		return errors.New("post db repository error: post is nil")
	}

	postResult, err := tx.ExecContext(ctx, "INSERT IGNORE INTO board_post(writer, title, content, like_cnt, time_to_read_minute) VALUES (?, ?, ?, ?, ?)", post.Writer, post.Title, post.Content, post.LikeCnt, post.TimeToReadMinute)
	if err != nil {
		return err
	}

	pid, err := postResult.LastInsertId()
	if err != nil {
		return err
	}

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

func (a *post) QueryBulkPost(ctx context.Context, tx *sqlx.Tx, opt *option.PostOption) ([]*entity.Post, error) {

	var (
		candPostList  []*entity.Post
		validPostList []*entity.Post
	)

	candPostList = make([]*entity.Post, 0)

	n, m, err := option.CalculateDBRange(opt.SizeRange)
	if err != nil {
		return nil, err
	}

	query := "SELECT bp.pid, bp.writer, bp.title, bp.content, bp.like_cnt, bp.time_to_read_minute, bp.create_at, COUNT(*) as views FROM board_post as bp LEFT OUTER JOIN board_views as bv ON bp.pid = bv.pid "
	if opt.QueryType == option.TitleAndContent {
		query += fmt.Sprintf("WHERE bp.title LIKE '%%%s%%' AND bp.content LIKE '%%%s%%' ", opt.Query, opt.Query)
	} else if opt.QueryType == option.Title {
		query += fmt.Sprintf("WHERE bp.title LIKE '%%%s%%' ", opt.Query)
	} else if opt.QueryType == option.Content {
		query += fmt.Sprintf("WHERE bp.content LIKE '%%%s%%' ", opt.Query)
	} else if opt.QueryType == option.Writer {
		query += fmt.Sprintf("WHERE bp.writer LIKE '%%%s%%' ", opt.Query)
	}

	query += fmt.Sprintf("GROUP BY bp.pid ORDER BY pid DESC LIMIT %d, %d", n, m)

	err = tx.SelectContext(ctx, &candPostList, query)
	if err != nil {
		return nil, err
	}

	validPostList = make([]*entity.Post, 0)

	for _, post := range candPostList {
		err = tx.SelectContext(ctx, &post.Tags, "SELECT value FROM board_tag WHERE board_tag.tid IN (SELECT board_post_tag.tid FROM board_post_tag WHERE board_post_tag.pid = ?)", post.Id)
		var suc = true
		for _, tag := range opt.Tags {
			var check = false
			for _, pTag := range post.Tags {
				if pTag == tag {
					check = true
				}
			}
			if check == false {
				suc = false
			}
		}
		if suc {
			validPostList = append(validPostList, post)
		}
	}

	return validPostList, nil
}
