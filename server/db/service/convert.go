package service

import (
	"github.com/dtc03012/me/db/entity"
	"github.com/dtc03012/me/protobuf/proto/entity/post"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func convertEntityPost(p *entity.Post) *post.Data {

	convertPost := &post.Data{
		Id:               p.Id,
		Title:            p.Title,
		Writer:           p.Writer,
		Content:          p.Content,
		Tags:             p.Tags,
		TimeToReadMinute: p.TimeToReadMinute,
		LikeCnt:          p.LikeCnt,
		IsNotice:         p.IsNotice,
		Views:            p.Views,
		CreateAt:         timestamppb.New(p.CreateAt),
	}

	return convertPost
}

func convertEntityPostList(postList []*entity.Post) []*post.Data {

	convertPostList := make([]*post.Data, 0, len(postList))

	for _, p := range postList {
		convertPostList = append(convertPostList, convertEntityPost(p))
	}

	return convertPostList
}

func convertEntityComment(c *entity.Comment) *post.Comment {

	convertComment := &post.Comment{
		PostId:   c.PostId,
		Writer:   c.Writer,
		Password: c.Comment,
		Comment:  c.Comment,
		LikeCnt:  c.LikeCnt,
	}

	return convertComment
}

func convertEntityCommentList(commentList []*entity.Comment) []*post.Comment {

	convertCommentList := make([]*post.Comment, 0, len(commentList))
	for _, c := range commentList {
		convertCommentList = append(convertCommentList, convertEntityComment(c))
	}

	return convertCommentList
}
