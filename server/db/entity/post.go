package entity

import "time"

type Post struct {
	Id               int32     `db:"pid"`
	Writer           string    `db:"writer"`
	Tags             []string  `db:"value"`
	Title            string    `db:"title"`
	Content          string    `db:"content"`
	LikeCnt          int32     `db:"like_cnt"'`
	Views            int32     `db:"views"`
	TimeToReadMinute int32     `db:"time_to_read_minute"`
	CreateAt         time.Time `db:"create_at"`
}
