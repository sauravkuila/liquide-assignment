package blog

import (
	"context"
	"liquide-assignment/pkg/dto"
	"strconv"

	"github.com/go-redis/redis/v8"
)

type postScore struct {
	PostID int64
	Score  float64
}

func (prs *blogSt) AddPost(ctx context.Context, post dto.Post) error {
	member := strconv.FormatInt(post.PostId, 10)
	// ZAdd to add a new member with initial score 0.

	_, err := prs.redisClient.ZAdd(ctx, prs.zsetKey, &redis.Z{
		Member: member,
		Score:  0,
	}).Result()
	return err
}

func (prs *blogSt) AddComment(ctx context.Context, comment dto.Comment) error {
	// member := strconv.FormatInt(comment.PostId, 10)

	//add to comment redis-hash
	return prs.updatePostScore(ctx, comment.PostId, 0.5)
}

func (prs *blogSt) AddVote(ctx context.Context, vote dto.Vote) error {

	if vote.VoteType == "upvote" {
		return prs.updatePostScore(ctx, vote.PostId, 1)
	}
	return prs.updatePostScore(ctx, vote.PostId, -1)
}

// updates the score for a given post by the provided delta.
// delta = +1 for an upvote, delta = -1 for a downvote, delta = +0.5 for a comment.
func (prs *blogSt) updatePostScore(ctx context.Context, postID int64, delta float64) error {
	member := strconv.FormatInt(postID, 10)

	_, err := prs.redisClient.ZIncrBy(ctx, prs.zsetKey, delta, member).Result()
	return err
}

// retrieves posts sorted in descending order by score.
// It supports pagination via offset and limit.
func (prs *blogSt) GetFeed(ctx context.Context, offset, limit int64) ([]postScore, error) {

	zSlice, err := prs.redisClient.ZRevRangeWithScores(ctx, prs.zsetKey, offset, offset+limit-1).Result()
	if err != nil {
		return nil, err
	}

	posts := make([]postScore, len(zSlice))
	for i, z := range zSlice {

		postID, err := strconv.ParseInt(z.Member.(string), 10, 64)
		if err != nil {
			return nil, err
		}
		posts[i] = postScore{
			PostID: postID,
			Score:  z.Score,
		}
	}
	return posts, nil
}
