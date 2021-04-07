VERSION = 0.1.0
TAG = $(VERSION)

GIT_COMMIT = $(shell git rev-parse HEAD)

GOLANGCI_CONTAINER=golangci/golangci-lint:v1.29-alpine
DATE= $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")

export DOCKER_BUILDKIT = 1

.PHONY: graylog-alert-exporter
graylog-alert-exporter:
	GO111MODULE=on CGO_ENABLED=0 go build -ldflags "-X main.version=$(VERSION) -X main.commit=$(GIT_COMMIT) -X main.date=$(DATE)" -o graylog-alert-exporter

.PHONY: lint
lint:
	docker run --rm -v $(shell pwd):go/src/git.opsta.io/graylog-alert-exporter -w go/src/git.opsta.io/graylog-alert-exporter $(GOLANGCI_CONTAINER) golangci-lint run

.PHONY: clean
clean:
	-rm graylog-alert-exporter
