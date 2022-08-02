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

	err := tx.Select(&post, "SELECT * FROM board_post WHERE pid = ?", pid)

	if err != nil {
		return nil, err
	}

	if len(post) == 0 {
		return nil, errors.New("post db error: there is no file id in db")
	}

	if len(post) > 1 {
		return nil, errors.New("post db error: duplicate file id. it is caused by server error")
	}

	err = tx.Select(&post[0].Tags, "SELECT value FROM board_tag WHERE board_tag.tid IN (SELECT board_post_tag.tid FROM board_post_tag WHERE board_post_tag.pid = ?)", pid)

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

	err = tx.Select(&posts, "SELECT * FROM board_post ORDER BY pid DESC LIMIT ?, ?", n, m)

	if err != nil {
		return nil, err
	}

	for _, post := range posts {
		err = tx.Select(&post.Tags, "SELECT value FROM board_tag WHERE board_tag.tid IN (SELECT board_post_tag.tid FROM board_post_tag WHERE board_post_tag.pid = ?)", post.Id)
	}

	if err != nil {
		return nil, err
	}

	return posts, nil
}

func (a *post) InsertPost(ctx context.Context, tx *sqlx.Tx, post *entity.Post, tags []string) error {
	if post == nil {
		return errors.New("post db error: post is nil")
	}

	postResult, err := tx.Exec("INSERT IGNORE INTO board_post(writer, title, content, time_to_read_minute) VALUES (?, ?, ?, ?)", post.Writer, post.Title, post.Content, post.TimeToReadMinute)
	if err != nil {
		return err
	}

	pid, err := postResult.LastInsertId()
	if err != nil {
		return err
	}

	for _, tag := range tags {
		tagResult, err := tx.Exec("INSERT IGNORE INTO board_tag(value) VALUES (?)", tag)
		if err != nil {
			return err
		}

		tid, err := tagResult.LastInsertId()
		if err != nil {
			return err
		}

		_, err = tx.Exec("INSERT IGNORE INTO board_post_tag(tid, pid) VALUES(?, ?)", tid, pid)
		if err != nil {
			return err
		}
	}

	return nil
}
