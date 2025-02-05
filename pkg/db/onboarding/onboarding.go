package onboarding

import (
	"liquide-assignment/pkg/dto"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type onboardingDb struct {
	dbObj *gorm.DB
}

type DbOnboardingtInterface interface {
	AddUser(*gin.Context, dto.DbUserDetail) (int64, error)
	GetUserByUsername(*gin.Context, string) (dto.DbUserDetail, error)
}

func NewOnboardingDbObject(db *gorm.DB) DbOnboardingtInterface {
	return &onboardingDb{
		dbObj: db,
	}
}
