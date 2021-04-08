# Exporter graylog alert

This exporter will receive alert events from graylog and create metrics about alert description for prometheus

[![Go Report Card](https://goreportcard.com/badge/github.com/StartloJ/graylog-alert-exporter)](https://goreportcard.com/report/github.com/StartloJ/graylog-alert-exporter)
[![GitHub tag](https://img.shields.io/github/tag/StartloJ/graylog-alert-exporter.svg)](https://github.com/StartloJ/graylog-alert-exporter/releases/latest)
[![GitHub](https://img.shields.io/github/license/StartloJ/graylog-alert-exporter)](https://github.com/StartloJ/graylog-alert-exporter/blob/main/LICENSE)

- [Exporter graylog alert](#exporter-graylog-alert)
  - [Installation](#installation)
  - [Usage](#usage)
  - [Route](#route)
  - [Config with Environment Variables and Flag](#config-with-environment-variables-and-flag)
  - [graylog spec](#graylog-spec)
## Installation

> go get github.com/StartloJ/graylog-alert-exporter

See [Release](https://github.com/StartloJ/graylog-alert-exporter/releases) page for maunal installation

## Usage
Download binary from release
```bash
# Run default parameter
$ cd to/path/of/file
$ ./graylog-alert-exporter
```
Example run with debug mode
```bash
$ ./graylog-alert-exporter --debug
```
Example full options
```bash
$ ./graylog-alert-exporter --listen '0.0.0.0:9889' \
  --path '/metrics' \
  --timeout 300 \
  --interval 5 \
  --debug \
  --caller
```


> All of above flag you can set with environment variables with `EXPORTER` prefix. example to enable dashboard you can run `EXPORTER_DASHBOARD=true graylog-alert-exporter`

## Route

1. / -> Provide home page will redirect to metrics.
1. /metrics -> Receive event from graylog with http POST method see api spec and Provision data metrics in http GET method for prometheus

## Config with Environment Variables and Flag

Name of env|Description|Type|Example
---|---|---|---
EXPORTER_LISTEN|IP and port that bind on host|string|0.0.0.0:9889
EXPORTER_PATH|Http sub-path provide metrics|string|/metrics
EXPORTER_TIMEOUT|Max time counter to resolved alert on metrics in second|number|60
EXPORTER_INTERVAL|Frequency of concurrent to check metrics,that should resolve |number|5
EXPORTER_DEBUG|Enable debugging log on console(stdout/stderr)|bool|false/true
EXPORTER_CALLER|Enable caller function to debugging|bool|false/true
EXPORTER_LABEL_FILE|Path to get yaml file for metrics label structure|string|example.yaml

> In explaination of exporter's flag you can see with `--help`
```txt
Usage:
      --caller          enable log method caller in code
      --debug           enable debug log
      --interval int    interval to check timeout (lower value consume more cpu) (default 5)
      --listen string   Host address to start service listen (default "0.0.0.0:9889")
      --path string     path for scape and push metrics (default "metrics")
      --timeout int     timeout of alert to make it resolved (default 60)
      --version         print version
      --label_file      Map labels config file to dynamic label in Prometheus metrics
```

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
