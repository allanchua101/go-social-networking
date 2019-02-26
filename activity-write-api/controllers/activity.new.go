package controllers

import (
	"github.com/gin-gonic/gin"
)

func HandleNewActivity(c *gin.Context) {
	// TODO: Validate parameters
	// TODO: Send Invalid Response
	// TODO: Offload queue message
	// TODO: Circuit break to second queue
	// TODO: Circuit break to a log store.
	// TODO: Trigger alarms about downtime + Send retry indicator to client side.
	RespondOK(c, gin.H{
		"payload": "Demo response",
	})
}
