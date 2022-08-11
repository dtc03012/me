package entity

import "time"

type Comment struct {
	Id              int32     `db:"cid"`
	PostId          int32     `db:"pid"`
	ParentCommentId int32     `db:"parent_cid"`
	IsExist         bool      `db:"is_exist"`
	Writer          string    `db:"writer"`
	Password        string    `db:"password"`
	Comment         string    `db:"comment"`
	LikeCnt         int32     `db:"like_cnt"`
	CreateAt        time.Time `db:"create_at"`
}
