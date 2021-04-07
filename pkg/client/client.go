package client

import (
	"graylog-alert-exporter/pkg/prom"
	"net/http"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

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

var AlertBody GraylogOutput

func UserServePayload(c *fiber.Ctx) error {
	err := c.BodyParser(&AlertBody)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"status":  "error",
			"message": err,
		})
	}
	alertCollect(&AlertBody)
	return c.Status(http.StatusCreated).JSON(&fiber.Map{
		"status":  "success",
		"message": "Created",
	})
}

type AlertMetrics struct {
	EventTitle       string
	EventDescription string
	EventTimeStamp   *time.Time
	EventSource      string
	EventPriority    int
}

func alertCollect(g *GraylogOutput) {
	newMetrics := AlertMetrics{
		EventTitle:       g.EventDefinitionTitle,
		EventDescription: g.EventDefinitionDescription,
		EventTimeStamp:   g.Event.Timestamp,
		EventSource:      g.Event.Source,
		EventPriority:    g.Event.Priority,
	}

	// factory := promauto.With(prom.Registry)
	// optEvent := prometheus.NewCounterVec(
	// 	prometheus.CounterOpts{
	// 		Name: "main_metrics",
	// 		Help: "Main metric",
	// 	},
	// 	[]string{"event_title", "event_description", "event_timestamp", "event_source", "event_priority"},
	// )
	// prometheus.MustRegister(optEvent)
	prom.OptEvent.WithLabelValues(
		newMetrics.EventTitle,
		newMetrics.EventDescription,
		newMetrics.EventSource,
		strconv.Itoa(newMetrics.EventPriority),
	).Inc()
}
