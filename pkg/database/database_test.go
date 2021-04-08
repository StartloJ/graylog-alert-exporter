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
			ID: utils.Hash(fmt.Sprintf("Test Title %d", i)),
			Data: map[string]string{
				"title":       fmt.Sprintf("Test Title %d", i),
				"description": "Test Desc",
				"source":      fmt.Sprintf("Test Source %d", i),
				"priority":    "critical",
			},
			Timeout: 60,
		})
	}

	alerts := database.GetAllAlerts()
	if len(alerts) != 10 {
		t.Error("Total log in database not equal 1")
	}
}
