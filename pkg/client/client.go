package client

import (
	"graylog-alert-exporter/pkg/prom"
	"net/http"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/prometheus/client_golang/prometheus"
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
	NewMetrics(&AlertBody)
	return nil
}

type AlertMetrics struct {
	EventID          string
	EventTitle       string
	EventDescription string
	EventTimeStamp   *time.Time
	EventSource      string
	EventPriority    int
}

func NewMetrics(g *GraylogOutput) {
	newMetrics := AlertMetrics{
		EventID:          g.EventDefinitionID,
		EventTitle:       g.EventDefinitionTitle,
		EventDescription: g.EventDefinitionDescription,
		EventTimeStamp:   g.Event.Timestamp,
		EventSource:      g.Event.Source,
		EventPriority:    g.Event.Priority,
	}
	metrics := prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "main_metrics",
			Help: "Main metric",
			ConstLabels: prometheus.Labels{
				"event_title":       newMetrics.EventTitle,
				"event_description": newMetrics.EventDescription,
				"event_timestamp":   newMetrics.EventTimeStamp.String(),
				"event_source":      newMetrics.EventSource,
				"event_priority":    strconv.Itoa(newMetrics.EventPriority),
			},
		},
	)
	prom.Registry.MustRegister(metrics)
}

func (c AlertMetrics) Describe(ch chan<- *prometheus.Desc) {
	ch <- prometheus.NewDesc("dummy", "dummy", nil, nil)
}
