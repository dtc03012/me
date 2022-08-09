package repository

import (
	"context"
	"errors"
	"github.com/dtc03012/me/db/entity"
	"github.com/dtc03012/me/db/option"
	"github.com/jmoiron/sqlx"
)

type post struct {
}

func (a *post) GetPost(ctx context.Context, tx *sqlx.Tx, pid int32) (*entity.Post, error) {

	var post []*entity.Post
	post = make([]*entity.Post, 0)

	err := tx.SelectContext(ctx, &post, "SELECT * FROM board_post WHERE pid = ?", pid)

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

	var posts []*entity.Post
	posts = make([]*entity.Post, 0)

	n, m, err := option.CalculateDBRange(opt.SizeRange)
	if err != nil {
		return nil, err
	}

	err = tx.SelectContext(ctx, &posts, "SELECT * FROM board_post ORDER BY pid DESC LIMIT ?, ?", n, m)

	if err != nil {
		return nil, err
	}

	for _, post := range posts {
		err = tx.SelectContext(ctx, &post.Tags, "SELECT value FROM board_tag WHERE board_tag.tid IN (SELECT board_post_tag.tid FROM board_post_tag WHERE board_post_tag.pid = ?)", post.Id)
	}

	if err != nil {
		return nil, err
	}

	return posts, nil
}

func (a *post) InsertPost(ctx context.Context, tx *sqlx.Tx, post *entity.Post, tags []string) error {
	if post == nil {
		return errors.New("post db repository error: post is nil")
	}

	postResult, err := tx.ExecContext(ctx, "INSERT IGNORE INTO board_post(writer, title, content, like_cnt, views, time_to_read_minute) VALUES (?, ?, ?, ?, ?, ?)", post.Writer, post.Title, post.Content, post.LikeCnt, post.Views, post.TimeToReadMinute)
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

func (a *post) GetViews(ctx context.Context, tx *sqlx.Tx, pid int32) (int, error) {

	var views []int32

	err := tx.SelectContext(ctx, &views, "SELECT views FROM board_post WHERE pid = ?", pid)
	if err != nil {
		return 0, err
	}

	if len(views) == 0 {
		return 0, errors.New("get views db repository error: there is no given pid in db")
	}

	return int(views[0]), nil
}

func (a *post) UpdateViews(ctx context.Context, tx *sqlx.Tx, views int32, pid int32) error {

	_, err := tx.ExecContext(ctx, "UPDATE board_post SET board_post.views = ? WHERE pid = ?", views, pid)
	if err != nil {
		return err
	}

	return nil
}

func (a *post) GetBulkComment(ctx context.Context, tx *sqlx.Tx, opt *option.CommentOption) ([]*entity.Comment, error) {

	var comments []*entity.Comment
	comments = make([]*entity.Comment, 0)

	n, m, err := option.CalculateDBRange(opt.SizeRange)
	if err != nil {
		return nil, err
	}

	err = tx.SelectContext(ctx, &comments, "SELECT * FROM board_comment WHERE pid = ? ORDER BY cid LIMIT ?, ?", opt.PostId, n, m)
	if err != nil {
		return nil, err
	}

	return comments, nil
}

func (a *post) InsertComment(ctx context.Context, tx *sqlx.Tx, comment *entity.Comment) error {

	if comment == nil {
		return errors.New("insert comment db repository error: comment is nil")
	}

	_, err := tx.ExecContext(ctx, "INSERT IGNORE INTO board_comment(pid, writer, password, comment, like_cnt) VALUES(?, ?, ?, ?, ?)",
		comment.PostId, comment.Writer, comment.Password, comment.Comment, comment.LikeCnt)
	if err != nil {
		return err
	}

	return nil
}

func (a *post) DeleteComment(ctx context.Context, tx *sqlx.Tx, postId int, commentId int) error {

	_, err := tx.ExecContext(ctx, "DELETE FROM board_comment WHERE pid = ? and cid = ?", postId, commentId)
	if err != nil {
		return err
	}

	return nil
}
