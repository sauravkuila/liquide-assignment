package post

import (
	"context"
	"liquide-assignment/pkg/dto"

	"gorm.io/gorm"
)

type postDb struct {
	dbObj *gorm.DB
}

type DbPostInterface interface {
	CreatePost(ctx context.Context, post dto.DbPost) (int64, error)
	UpdatePost(ctx context.Context, post dto.DbPost) (int64, error)
	DeletePost(ctx context.Context, postId int64, userId int64) error
	GetPost(ctx context.Context, postId int64) (dto.DbPostInfo, error)
	GetUserPosts(ctx context.Context, userId int64, limit, offset int) ([]dto.PostInfo, int64, error)
}

func NewPostDbObject(db *gorm.DB) DbPostInterface {
	return &postDb{
		dbObj: db,
	}
}
