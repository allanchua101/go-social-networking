package controllers

import (
	"activity-write-api/messaging"
	"activity-write-api/models"
	"activity-write-api/validators"

	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/twinj/uuid"
	"os"
)

// HandleNewActivityRequest handles HTTP request for
// pushing new activity data into the system.
func HandleNewActivityRequest(c *gin.Context) {
	input := new(models.Activity)
	err := c.BindJSON(input)

	if err != nil {
		SendBadRequest(c, gin.H{
			"payload": "Request cannot be parsed.",
		})
		return
	}
	
	validationIssues := validators.ValidateNewActivity(input)

	if len(validationIssues) > 0 {
		SendUnprocessableEntity(c, gin.H {
			"payload": validationIssues,
		})
		return
	}
	
	input.ID = uuid.NewV4().String()

	SendOK(c, gin.H{
		"payload": input.ID,
	})
	
	writeQueueName := os.Getenv("WRITE_API_QUEUE_NAME")
	serializedData, serializeErr := json.Marshal(input)

	if serializeErr == nil {
		publishErr := messaging.PublishEvent(writeQueueName, string(serializedData))

		fmt.Println(publishErr)
		// TODO: Circuit break to a log store.
		// TODO: Trigger alarms about downtime + Send retry indicator to client side.
	}
}
