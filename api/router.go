package api

import (
	"liquide-assignment/pkg/service"

	"github.com/gin-gonic/gin"
)

func getRouter(obj service.ServiceGroupLayer) *gin.Engine {
	router := gin.New()
	router.Use(gin.Recovery())

	// Health check API can be used for the Kubernetes pod health
	router.GET("/health", obj.Health)

	//cred APIs
	onboardingGroup := router.Group("onboarding")
	{
		onboardingGroup.POST("signup", obj.GetOnboardingService().UserSignup) //signup as user or admin
		onboardingGroup.POST("login", obj.GetOnboardingService().UserLogin)   //login for user / admin
	}

	return router
}
