package db

import (
	"liquide-assignment/pkg/db/onboarding"

	"gorm.io/gorm"
)

type dbService struct {
	onboarding.DbOnboardingtInterface
}

type DBLayer interface {
	onboarding.DbOnboardingtInterface
}

func NewDBObject(psqlDB *gorm.DB) DBLayer {
	temp := &dbService{
		onboarding.NewOnboardingDbObject(psqlDB),
	}
	return temp
}
