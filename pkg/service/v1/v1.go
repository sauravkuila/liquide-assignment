package v1

import (
	"liquide-assignment/pkg/db"
	"liquide-assignment/pkg/service/v1/comment"
	"liquide-assignment/pkg/service/v1/post"
	"liquide-assignment/pkg/service/v1/vote"
)

type v1ServiceObj struct {
	post.PostServiceInterface
	comment.CommentServiceInterface
	vote.VoteServiceInterface
}

type V1ServiceInterface interface {
	post.PostServiceInterface
	comment.CommentServiceInterface
	vote.VoteServiceInterface
}

func NewServiceObject(db db.DBLayer) V1ServiceInterface {
	return &v1ServiceObj{
		PostServiceInterface:    post.NewPostService(db),
		CommentServiceInterface: comment.NewCommentService(db),
		VoteServiceInterface:    vote.NewVoteService(db),
	}
}
