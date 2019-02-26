package main

import (
	"activity-write-api/controllers"
	"github.com/gin-gonic/gin"
)

func registerRoutes(router *gin.Engine) {
	activityRoutes := router.Group("/activity")
	{
		activityRoutes.POST("/new", controllers.HandleNewActivity)
	}
}
