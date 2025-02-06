package comment

import (
	"context"
	"liquide-assignment/pkg/dto"

	"gorm.io/gorm"
)

type commentDb struct {
	dbObj *gorm.DB
}

type DbCommentInterface interface {
	AddComment(ctx context.Context, comment dto.DbComment) (int64, error)
	AddReply(ctx context.Context, comment dto.DbComment) (int64, error)
	UpdateComment(ctx context.Context, comment dto.DbComment) (int64, error)
	DeleteComment(ctx context.Context, commentId int64, userId int64) error
}

func NewCommentDbObject(db *gorm.DB) DbCommentInterface {
	return &commentDb{
		dbObj: db,
	}
}
