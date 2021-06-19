package database_test

import (
	"fmt"
	"graylog-alert-exporter/internal/utils"
	"graylog-alert-exporter/pkg/database"
	"reflect"
	"testing"
)

func TestAddAlert(t *testing.T) {
	database.Init()
	alert := database.Alert{
		ID: utils.Hash("Test Title"),
		Data: map[string]string{
			"title":       "test title",
			"description": "test description",
			"source":      "test source",
			"priority":    "test",
		},
		Timeout: 60,
	}

	if err := database.InsertAlert(alert); err != nil {
		t.Error("Error to insert record into database: ", err)
	}

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
		if err := database.InsertAlert(database.Alert{
			ID: utils.Hash(fmt.Sprintf("Test Title %d", i)),
			Data: map[string]string{
				"title":       fmt.Sprintf("Test Title %d", i),
				"description": "Test Desc",
				"source":      fmt.Sprintf("Test Source %d", i),
				"priority":    "critical",
			},
			Timeout: 60,
		}); err != nil {
			t.Error("Error to insert record into database: ", err)
		}
	}

	alerts := database.GetAllAlerts()
	if len(alerts) != 10 {
		t.Error("Total log in database not equal 1")
	}
}

func TestMissingRequiredJsonField(t *testing.T) {
	database.Init()

	if err := database.InsertAlert(database.Alert{
		ID: "",
		Data: map[string]string{
			"title":       "Test Title",
			"description": "Test Desc",
			"source":      "Test Source",
			"priority":    "critical",
		},
		Timeout: 60,
	}); err == nil {
		t.Error("This should error because ID is empty")
	}
}
