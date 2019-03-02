package persistence

import (
	"os"
	"log"
	"encoding/json"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	
	"activity-write-daemon/models"
)

// PushEvent method is used for storing activities / entries
// to the event store.
func PushEvent(activity *models.Activity) bool {
	db, err := gorm.Open("postgres", os.Getenv("EVENT_STORE_STR"))

	if err != nil {
		log.Printf("[Write Daemon Error]%s\n", err)
		return false
	}
	
	data, err := json.Marshal(activity)

	if err != nil {
		log.Printf("[Write Daemon Error]:%s\n", err)
		return false
	}

	db.Exec("INSERT INTO events (\"id\", \"data\", \"aggregate\") VALUES (?, ?, ?)", activity.ID, data, activity.Verb)
	defer db.Close()

	return true
}
