package main

import (
	"github.com/gin-gonic/gin"
)

func registerRoutes(router *gin.Engine) {
	activityRoutes := router.Group("/activity")
	{
		activityRoutes.POST("/new", handleNewActivity)
	}
}
