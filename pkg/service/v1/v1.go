package v1

import (
	"liquide-assignment/pkg/db"
	"liquide-assignment/pkg/service/v1/comment"
	"liquide-assignment/pkg/service/v1/feed"
	"liquide-assignment/pkg/service/v1/post"
	"liquide-assignment/pkg/service/v1/vote"

	"github.com/go-redis/redis/v8"
)

type v1ServiceObj struct {
	post.PostServiceInterface
	comment.CommentServiceInterface
	vote.VoteServiceInterface
	feed.FeedServiceInterface
}

type V1ServiceInterface interface {
	post.PostServiceInterface
	comment.CommentServiceInterface
	vote.VoteServiceInterface
	feed.FeedServiceInterface
}

func NewServiceObject(db db.DBLayer, redisConn *redis.Client) V1ServiceInterface {
	return &v1ServiceObj{
		post.NewPostService(db, redisConn),
		comment.NewCommentService(db, redisConn),
		vote.NewVoteService(db, redisConn),
		feed.NewFeedService(db, redisConn),
	}
}
