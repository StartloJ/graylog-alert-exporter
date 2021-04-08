VERSION = 0.1.0
TAG = $(VERSION)
GIT_COMMIT = $(shell git rev-parse HEAD)
DATE= $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
GOLANGCI_CONTAINER=golangci/golangci-lint:v1.39-alpine

export DOCKER_BUILDKIT = 1

.PHONY: graylog-alert-exporter
build:
	GO111MODULE=on CGO_ENABLED=0 go build -ldflags "-X main.version=$(VERSION) -X main.commit=$(GIT_COMMIT) -X main.date=$(DATE)" -o graylog-alert-exporter

docker:
	docker build -t opsta/grayexporter:$(TAG) --build-arg=COMMIT=$(git rev-parse --short HEAD) .

.PHONY: lint
lint:
	docker run --rm -v $(shell pwd):/go/src/git.opsta.io/graylog-alert-exporter -w /go/src/git.opsta.io/graylog-alert-exporter $(GOLANGCI_CONTAINER) golangci-lint run

.PHONY: clean
clean:
	-rm graylog-alert-exporter
