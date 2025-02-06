package feed

import (
	"liquide-assignment/pkg/blog"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (obj *feedService) GetFeed(c *gin.Context) {

	// Get the feed for a user where posts are sorted by popularity
	// popularty is calculated by the number of likes and comments on the post
	// score of a post can be determined by this formula
	// score=upvotes−downvotes+comment_weight×(number of comments)−time decay
	blogObj := blog.NewBlogObject(obj.redisObj)

	posts, err := blogObj.GetFeed(c, 0, 10)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch feed"})
		return
	}
	// Return the feed
	c.JSON(http.StatusOK, posts)
}
