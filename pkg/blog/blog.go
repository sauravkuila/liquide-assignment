package blog

import (
	"context"
	"liquide-assignment/pkg/dto"

	"github.com/go-redis/redis/v8"
)

type blogSt struct {
	redisClient *redis.Client
	zsetKey     string
}

var blogobj *blogSt = nil

type BlogInterface interface {
	AddPost(ctx context.Context, post dto.Post) error
	AddComment(ctx context.Context, comment dto.Comment) error
	AddVote(ctx context.Context, comment dto.Vote) error
	GetFeed(ctx context.Context, offset, limit int64) ([]postScore, error)
}

func NewBlogObject(redisClient *redis.Client) BlogInterface {
	if blogobj == nil {
		blogobj = &blogSt{
			redisClient: redisClient,
			zsetKey:     "liquide-blog",
		}
	}
	return blogobj
}
