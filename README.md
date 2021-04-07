# Exporter graylog alert
This exporter will recieve alert events from graylog and create metrics about alert description

## Route
1. / -> Provide home page will redirect to metrics.
1. /store -> Recieve event from graylog with http post method see api spec
1. /metrics -> Provision data metrics for prometheus

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