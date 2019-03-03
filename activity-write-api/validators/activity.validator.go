package validators

import (
	"activity-write-api/models"
)

const followVerb string = "follow"
const postVerb string = "post"
const shareVerb string = "share"
const likeVerb string = "like"
const unfollowVerb string = "unfollow"
const activityActorRequiredMessage = "Actor is required."
const activityVerbRequiredMessage = "Verb is required."
const activityObjectRequiredMessage = "Object is required for this verb."
const activityTargetRequiredMessage = "Target is required for this verb."
const invalidVerbMessage = "Provided verb is invalid."

func isEmpty(input string) bool {
	return input == ""
}

func isValidVerb(input string) bool {
	set := [5]string{followVerb, postVerb, shareVerb, likeVerb, unfollowVerb}

	for _, item := range set {
			if item == input {
					return true
			}
	}
	return false
}

// ValidateNewActivity validates if a new activity
// is a valid input or not.
func ValidateNewActivity(activity *models.Activity) []string  {
	var output []string

	if isEmpty(activity.Actor) {
		output = append(output, activityActorRequiredMessage)
	}

	if !isValidVerb(activity.Verb) {
		output = append(output, invalidVerbMessage)
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