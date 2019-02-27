package controllers

import (
	"activity-write-api/models"
	"github.com/gin-gonic/gin"
)

const followVerb string = "follow"
const postVerb string = "post"
const activityActorRequiredMessage = "Actor is required."
const activityVerbRequiredMessage = "Verb is required."
const activityObjectRequiredMessage = "Object is required for this verb."
const activityTargetRequiredMessage = "Target is required for this verb."
const emptyString = ""

func isEmpty(input string) bool {
	return input == emptyString
}

func validateNewActivityRequest(activity *models.Activity) []string  {
	var output []string

	if isEmpty(activity.Actor) {
		output = append(output, activityActorRequiredMessage)
	}

	if isEmpty(activity.Verb) {
		output = append(output, activityVerbRequiredMessage)
	}

	if activity.Verb != followVerb  && isEmpty(activity.Object) {
		output = append(output, activityObjectRequiredMessage)
	}

	if activity.Verb != postVerb && isEmpty(activity.Target) {
		output = append(output, activityTargetRequiredMessage)
	}

	return output
}

func HandleNewActivity(c *gin.Context) {
	input := new(models.Activity)
	err := c.BindJSON(input)

	if err != nil {
		SendBadRequest(c, gin.H{
			"payload": "Request cannot be parsed.",
		})
	} else {
		validationIssues := validateNewActivityRequest(input)

		if len(validationIssues) > 0 {
			SendUnprocessableEntity(c, gin.H {
				"payload": validationIssues,
			})
		} else {
			// TODO: Offload queue message
			// TODO: Circuit break to second queue
			// TODO: Circuit break to a log store.
			// TODO: Trigger alarms about downtime + Send retry indicator to client side.
			SendOK(c, gin.H{
				"payload": "IT Seems to be OK",
			})
		}
	}
}
