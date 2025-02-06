package service

import (
	"liquide-assignment/pkg/db"
	"liquide-assignment/pkg/service/onboarding"
	v1 "liquide-assignment/pkg/service/v1"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

type service struct {
	onboardingService onboarding.OnboardingInterface
	v1Service         v1.V1ServiceInterface
}

type ServiceGroupLayer interface {
	Health(*gin.Context)
	GetOnboardingService() onboarding.OnboardingInterface
	GetV1Service() v1.V1ServiceInterface
}

func NewServiceGroupObject(db db.DBLayer, redisConn *redis.Client) ServiceGroupLayer {
	return &service{
		onboardingService: onboarding.NewOnboardingService(db),
		v1Service:         v1.NewServiceObject(db, redisConn),
	}
}

func (s *service) Health(c *gin.Context) {
	c.String(http.StatusOK, "I am Healthy")
}

func (s *service) GetOnboardingService() onboarding.OnboardingInterface {
	return s.onboardingService
}

func (s *service) GetV1Service() v1.V1ServiceInterface {
	return s.v1Service
}
