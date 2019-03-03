package parsers

import (
	"encoding/json"

	"my-feed-daemon/models"
)

// ParseActivityJson method is used for parsing JSON string
// to Activity struct instance.
func ParseActivityJson(jsonString []byte) (*models.Activity, error) {
	var activity *models.Activity
	activityData := []byte(jsonString)
	err := json.Unmarshal(activityData, &activity)

	return activity, err
}