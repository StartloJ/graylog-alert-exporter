# Exporter graylog alert

This exporter will receive alert events from graylog and create metrics about alert description for prometheus

[![Go Report Card](https://goreportcard.com/badge/github.com/StartloJ/graylog-alert-exporter)](https://goreportcard.com/report/github.com/StartloJ/graylog-alert-exporter)
[![GitHub tag](https://img.shields.io/github/tag/StartloJ/graylog-alert-exporter.svg)](https://github.com/StartloJ/graylog-alert-exporter/releases/latest)
[![GitHub](https://img.shields.io/github/license/StartloJ/graylog-alert-exporter)](https://github.com/StartloJ/graylog-alert-exporter/blob/main/LICENSE)

## Architecture Diagram

<!-- TODO: Add Architecture Diagram -->

## Table of Contents

- [Exporter graylog alert](#exporter-graylog-alert)
  - [Architecture Diagram](#architecture-diagram)
  - [Table of Contents](#table-of-contents)
  - [Installations](#installations)
    - [Go](#go)
    - [Prebuild binary](#prebuild-binary)
    - [Docker](#docker)
  - [Usages](#usages)
    - [Example with default config](#example-with-default-config)
    - [Use with dynamic labels](#use-with-dynamic-labels)
  - [Routes](#routes)
  - [Config with Environment Variables and Flags](#config-with-environment-variables-and-flags)
  - [How to use API](#how-to-use-api)
    - [For example use case, you start app in localhost and use default port.](#for-example-use-case-you-start-app-in-localhost-and-use-default-port)
  - [graylog spec](#graylog-spec)

## Installations

### Go

> go get github.com/StartloJ/graylog-alert-exporter

### Prebuild binary

```command
wget https://github.com/StartloJ/graylog-alert-exporter/releases/download/v0.1.0/graylog-alert-exporter_0.1.0_Linux_x86_64.tar.gz
tar -xvf graylog-alert-exporter_0.1.0_Linux_x86_64.tar.gz
mv graylog-alert-exporter /usr/local/bin/
```

See [Release](https://github.com/StartloJ/graylog-alert-exporter/releases) page for more version

### Docker

```command
docker run -d --name graylog-alert-exporter -p 9889:9889 StartloJ/graylog-alert-exporter
```

## Usages

### Example with default config

```command
graylog-alert-exporter --listen '0.0.0.0:9889' \
--path '/metrics' \
--timeout 120 \
--interval 5
```

> All of above flag you can set with environment variables with `EXPORTER` prefix. example to enable dashboard you can run `EXPORTER_DASHBOARD=true graylog-alert-exporter`

### Use with dynamic labels

```command
graylog-alert-exporter -f example-labels.yaml
```

see [this](example-labels.yaml) for example config file

## Routes

1. / -> Provide home page will redirect to metrics.
1. /metrics -> Receive event from graylog with http POST method see api spec and Provision data metrics in http GET method for prometheus

## Config with Environment Variables and Flags

Name of env|Description|Type|Example
---|---|---|---
EXPORTER_LISTEN|IP and port that bind on host|string|0.0.0.0:9889
EXPORTER_PATH|Http sub-path provide metrics|string|/metrics
EXPORTER_TIMEOUT|Max time counter to resolved alert on metrics in second|number|60
EXPORTER_INTERVAL|Frequency of concurrent to check metrics,that should resolve |number|5
EXPORTER_DEBUG|Enable debugging log on console(stdout/stderr)|bool|false/true
EXPORTER_CALLER|Enable caller function to debugging|bool|false/true
EXPORTER_LABEL_FILE|Path to get yaml file for metrics label structure|string|example.yaml
EXPORTER_DASHBOARD|enable web application to control alert metrics in GUI|bool|false/true

> In explaination of exporter's flag you can see with `--help`

```txt
Usage:
      --dashboard           enable web application to control alert metrics in GUI
      --debug               enable debug log
  -i, --interval int        interval to check timeout (lower value consume more cpu) in second (default 5)
  -f, --label_file string   Map labels config file to dynamic label in Prometheus metrics (default "labels.yaml")
  -l, --listen string       Host address to start service listen (default "0.0.0.0:9889")
  -p, --path string         path for scape and push metrics (default "metrics")
  -t, --timeout int         timeout of alert to make it resolved in second (default 60)
  -v, --version             print version
```

## How to use API
With **newly update** feature for support to manage metrics with `REST API`.  
You can **update**, **get** and **delete** metrics on the fly. You see all api path  
under `/api/alert(s)`.
- **Get method**: `/api/alerts` to get all metrics.
- **Delete method**: `/api/alert/<id>` to remove metrics with hash id.
- **Post method**: `/api/alert` to add new metrics with json structure:
  - ID: sha256 text
  - Timeout: integer
  - Data: map of json data

  
For example data in following code:
```json
{
  "ID": "8ff1d293923dde8fa49a3227b4a4b4faad94f73cd1c41c60e712d4f1e421a788",
  "Timeout": 30,
  "Data": {
    "statuscode":"500",
    "namespace":"router",
    "app_name":"nginx-tester",
    "some_info":"kitty",
    "zone":"OPSTA"
  }
}
```

### For example use case, you start app in localhost and use default port.
- Insert metrics with `cURL`. Before you send api, you will create hash string from alert title.

```console
$ echo "Found error 503 in application" | sha256sum
$ curl --request POST 'http://127.0.0.1:9889/api/alert' \
--header 'Content-Type: application/json' \
--data '{
  "ID": "1d9b3a1b3710f0d28cf2efd82fafa480df453de2c75458945639f7b4f3dd32cd",
  "Timeout": 30,
  "Data": {
    "statuscode":"500",
    "namespace":"router",
    "app_name":"nginx-tester",
    "some_info":"kitty",
    "zone":"OPSTA"
  }
}'
```
- Get all metrics in current database.

```console
$ curl --request GET 'http://127.0.0.1:9889/api/alerts'
```
- Delete inserted metrics

```console
$ MSG_HASH='1d9b3a1b3710f0d28cf2efd82fafa480df453de2c75458945639f7b4f3dd32cd'
$ curl --request DELETE http://127.0.0.1:9889/api/alert/$MSG_HASH
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
