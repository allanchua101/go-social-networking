package validators

import (
	"activity-write-api/models"
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

// ValidateNewActivity validates if a new activity
// is a valid input or not.
func ValidateNewActivity(activity *models.Activity) []string  {
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