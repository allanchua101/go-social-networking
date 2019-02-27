package controllers

import (
	"os"
	"fmt"
	"activity-write-api/models"
	"activity-write-api/validators"
	"activity-write-api/emitters"
	"github.com/gin-gonic/gin"
	"github.com/twinj/uuid"
	"encoding/json"
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
	} else {
		validationIssues := validators.ValidateNewActivity(input)

		if len(validationIssues) > 0 {
			SendUnprocessableEntity(c, gin.H {
				"payload": validationIssues,
			})
		} else {
			input.ID = uuid.NewV4().String()

			SendOK(c, gin.H{
				"payload": input.ID,
			})
			
			writeQueueName := os.Getenv("WRITE_API_QUEUE_NAME")
			serializedData, serializeErr := json.Marshal(input)

			fmt.Println("%s", writeQueueName)
			fmt.Println("%s", string(serializedData))

			if serializeErr == nil {
				publishErr := emitters.PublishEvent(writeQueueName, string(serializedData))

				fmt.Println("%s", publishErr)
				// TODO: Circuit break to a log store.
				// TODO: Trigger alarms about downtime + Send retry indicator to client side.
			}
		}
	}
}
