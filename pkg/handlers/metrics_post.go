package handlers

import (
	"graylog-alert-exporter/internal/utils"
	"graylog-alert-exporter/pkg/database"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

// GraylogOutput is json model send from graylog webhook
type GraylogOutput struct {
	EventDefinitionID          string `json:"event_definition_id"`
	EventDefinitionType        string `json:"event_definition_type"`
	EventDefinitionTitle       string `json:"event_definition_title"`
	EventDefinitionDescription string `json:"event_definition_description"`
	JobDefinitionID            string `json:"job_definition_id"`
	JobTriggerID               string `json:"job_trigger_id"`
	Event                      struct {
		ID                  string        `json:"id"`
		EventDefinitionType string        `json:"event_definition_type"`
		EventDefinitionID   string        `json:"event_definition_id"`
		OriginContext       string        `json:"origin_context"`
		Timestamp           *time.Time    `json:"timestamp"`
		TimestampProcessing *time.Time    `json:"timestamp_processing"`
		TimerangeStart      *time.Time    `json:"timerange_start,omitempty"`
		TimerangeEnd        *time.Time    `json:"timerange_end,omitempty"`
		Streams             []interface{} `json:"streams"`
		SourceStreams       []string      `json:"source_streams"`
		Message             string        `json:"message"`
		Source              string        `json:"source"`
		KeyTuple            []interface{} `json:"key_tuple"`
		Key                 string        `json:"key"`
		Priority            int           `json:"priority"`
		Alert               bool          `json:"alert"`
		Fields              struct{}      `json:"fields"`
	} `json:"event"`
	Backlog []struct {
		Index     string                 `json:"index"`
		Message   string                 `json:"message"`
		Timestamp *time.Time             `json:"timestamp"`
		Source    string                 `json:"source"`
		StreamIds []string               `json:"stream_ids"`
		ID        string                 `json:"id"`
		Fields    map[string]interface{} `json:"fields"`
	} `json:"backlog"`
}

// ExtractAlertMetrics return Alert from GraylogOutput
func (g GraylogOutput) ExtractAlertMetrics() (*database.Alert, error) {
	// Default labels
	severity, err := utils.RePriority(g.Event.Priority)
	if err != nil {
		return nil, err
	}
	alert := database.Alert{
		ID:      utils.Hash(g.EventDefinitionTitle),
		Timeout: viper.GetInt("timeout"),
		Data: map[string]string{
			"title":       g.EventDefinitionTitle,
			"description": g.EventDefinitionDescription,
			"source":      g.Event.Source,
			"priority":    severity,
		},
	}

	// Dynamic labels
	for k, v := range viper.GetStringMapString("labels") {
		alert.Data[k] = utils.GetValueFromJSON(v, g).(string)
	}

	return &alert, nil
}

// GetGraylogOutputHandler is handler for store graylog alert payload
func GetGraylogOutputHandler(c *fiber.Ctx) error {
	g := GraylogOutput{}
	err := c.BodyParser(&g)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "error",
			"message": err,
		})
	}

	alert, _ := g.ExtractAlertMetrics()
	database.InsertAlert(*alert)

	return c.Status(fiber.StatusCreated).JSON(&fiber.Map{
		"status":  "success",
		"message": "created",
	})
}
