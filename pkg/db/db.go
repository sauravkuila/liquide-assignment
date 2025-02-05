package db

import (
	"liquide-assignment/pkg/db/comment"
	"liquide-assignment/pkg/db/onboarding"
	"liquide-assignment/pkg/db/post"
	"liquide-assignment/pkg/db/vote"

	"gorm.io/gorm"
)

type dbService struct {
	onboarding.DbOnboardingtInterface
	post.DbPostInterface
	comment.DbCommentInterface
	vote.DbVoteInterface
}

type DBLayer interface {
	onboarding.DbOnboardingtInterface
	post.DbPostInterface
	comment.DbCommentInterface
	vote.DbVoteInterface
}

func NewDBObject(psqlDB *gorm.DB) DBLayer {
	temp := &dbService{
		onboarding.NewOnboardingDbObject(psqlDB),
		post.NewPostDbObject(psqlDB),
		comment.NewCommentDbObject(psqlDB),
		vote.NewVoteDbObject(psqlDB),
	}
	return temp
}
