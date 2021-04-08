# Exporter graylog alert

This exporter will receive alert events from graylog and create metrics about alert description for prometheus

[![Go Report Card](https://goreportcard.com/badge/github.com/StartloJ/graylog-alert-exporter)](https://goreportcard.com/report/github.com/StartloJ/graylog-alert-exporter)
[![GitHub tag](https://img.shields.io/github/tag/StartloJ/graylog-alert-exporter.svg)](https://github.com/StartloJ/graylog-alert-exporter/releases/latest)
[![GitHub](https://img.shields.io/github/license/StartloJ/graylog-alert-exporter)](https://github.com/StartloJ/graylog-alert-exporter/blob/main/LICENSE)

## Installation

> go get github.com/StartloJ/graylog-alert-exporter

See [Release](https://github.com/StartloJ/graylog-alert-exporter/releases) page for maunal installation

## Usage

```txt
Usage:
      --caller          enable log method caller in code
      --dashboard       enable dashboard web service listen at path dashboard
      --debug           enable debug log
      --interval int    interval to check timeout (lower value consume more cpu) (default 5)
      --listen string   Host address to start service listen (default "0.0.0.0:9889")
      --path string     path for scape and push metrics (default "metrics")
      --timeout int     timeout of alert to make it resolved (default 60)
      --version         print version
```

> All of above flag you can set with environment variables with `APP` prefix. example to enable dashboard you can run `APP_DASHBOARD=true graylog-alert-exporter`

## Route

1. / -> Provide home page will redirect to metrics.
2. /metrics -> Receive event from graylog with http POST method see api spec and Provision data metrics in http GET method for prometheus
3. /dashboard -> Basic host resources monitoring

## graylog spec

| Title                      | Type        | Example | required |
|----------------------------|-------------|---------|----------|
| EventDefinitionID          | string      |         | yes      |
| EventDefinitionType        | string      |         | yes      |
| EventDefinitionTitle       | string      |         | yes      |
| EventDefinitionDescription | string      |         | yes      |
| JobDefinitionID            | string      |         | yes      |
| JobTriggerID               | string      |         | yes      |
| Event                      | map         |         | yes      |
| Backlog                    | list of map |         | yes      |
