package vote

import (
	"context"
	"liquide-assignment/pkg/dto"

	"gorm.io/gorm"
)

type voteDb struct {
	dbObj *gorm.DB
}

type DbVoteInterface interface {
	AddVote(ctx context.Context, vote dto.DbVote) (int64, error)
	UpdateVote(ctx context.Context, vote dto.DbVote) (int64, error)
	UpsertVote(ctx context.Context, vote dto.DbVote) (int64, error)
}

func NewVoteDbObject(db *gorm.DB) DbVoteInterface {
	return &voteDb{
		dbObj: db,
	}
}
