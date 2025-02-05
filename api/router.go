package api

import (
	"liquide-assignment/pkg/auth"
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

	router.Use(auth.AuthMiddleware())

	//v1 APIs
	v1Group := router.Group("v1")
	{
		//enhancement: depending on scale, each group can become a microservice pointing to individual tables or databases if needed

		//post group
		postGroup := v1Group.Group("post")
		{
			postGroup.POST("", obj.GetV1Service().CreatePost)      //create post
			postGroup.PUT(":id", obj.GetV1Service().EditPost)      //edit a post
			postGroup.DELETE(":id", obj.GetV1Service().DeletePost) //delete a post
			postGroup.GET(":id", obj.GetV1Service().GetPost)       //fetch a post info
			postGroup.GET("", obj.GetV1Service().GetAllPosts)      //fetch all posts
		}

		//comment group
		commentGroup := v1Group.Group("comment")
		{
			commentGroup.POST("", obj.GetV1Service().CreateComment)       //create comment
			commentGroup.PUT(":id", obj.GetV1Service().EditComment)       //edit a comment
			commentGroup.DELETE(":id", obj.GetV1Service().DeleteComment)  //delete a comment
			commentGroup.GET(":id", obj.GetV1Service().GetComment)        //fetch a comment info
			commentGroup.GET(":postid", obj.GetV1Service().GetAllComment) //fetch all comments
		}

		//vote group
		voteGroup := v1Group.Group("vote")
		{
			voteGroup.POST("/up", obj.GetV1Service().UpVote)     //upvote
			voteGroup.POST("/down", obj.GetV1Service().DownVote) //downvote
		}

	}

	return router
}
