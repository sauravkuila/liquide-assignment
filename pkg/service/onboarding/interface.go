package onboarding

import (
	"liquide-assignment/pkg/db"

	"github.com/gin-gonic/gin"
)

type onboardingService struct {
	dbObj db.DBLayer
}

type OnboardingInterface interface {
	UserSignup(*gin.Context)
	UserLogin(*gin.Context)
}

func NewOnboardingService(db db.DBLayer) OnboardingInterface {
	return &onboardingService{
		dbObj: db,
	}
}
