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
		ID:          fmt.Sprintf("%x", sha256.Sum256([]byte("Test Title"+"Test Desc"))),
		Title:       "Test Title",
		Description: "Test Desc",
		Source:      "Test source",
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
