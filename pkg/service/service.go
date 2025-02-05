package service

import (
	"liquide-assignment/pkg/db"
	"liquide-assignment/pkg/service/onboarding"
	"net/http"

	"github.com/gin-gonic/gin"
)

type service struct {
	onboardingService onboarding.OnboardingInterface
}

type ServiceGroupLayer interface {
	Health(*gin.Context)
	GetOnboardingService() onboarding.OnboardingInterface
}

func NewServiceGroupObject(db db.DBLayer) ServiceGroupLayer {
	return &service{
		onboardingService: onboarding.NewOnboardingService(db),
	}
}

func (s *service) Health(c *gin.Context) {
	c.String(http.StatusOK, "I am Healthy")
}

func (s *service) GetOnboardingService() onboarding.OnboardingInterface {
	return s.onboardingService
}
