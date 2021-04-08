package database_test

import (
	"crypto/sha256"
	"fmt"
	"graylog-alert-exporter/pkg/database"
	"reflect"
	"testing"
)

func TestAddAlert(t *testing.T) {
	database.Init()
	alert := database.Alert{
		ID:          fmt.Sprintf("%x", sha256.Sum256([]byte("Test Title"+"Test Source"))),
		Title:       "Test Title",
		Description: "Test Desc",
		Source:      "Test Source",
		Priority:    3,
		Timeout:     60,
	}
	database.InsertAlert(alert)

	alerts := database.GetAllAlerts()
	if len(alerts) != 1 {
		t.Error("Total log in database not equal 1")
	}

	if !reflect.DeepEqual(alert, alerts[0]) {
		t.Error("Alert struct not equal struct get from database")
	}
}

func TestMultiAddAlertAndGetAll(t *testing.T) {
	database.Init()
	for i := 0; i < 10; i++ {
		database.InsertAlert(database.Alert{
			ID:          fmt.Sprintf("%x", sha256.Sum256([]byte(fmt.Sprintf("Test Title %d", i)+fmt.Sprintf("Test Source %d", i)))),
			Title:       fmt.Sprintf("Test Title %d", i),
			Description: "Test Desc",
			Source:      fmt.Sprintf("Test Source %d", i),
			Priority:    3,
			Timeout:     60,
		})
	}

	alerts := database.GetAllAlerts()
	if len(alerts) != 10 {
		t.Error("Total log in database not equal 1")
	}
}
