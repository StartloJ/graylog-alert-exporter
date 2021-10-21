FROM golang:1.16.3-alpine as builder
WORKDIR /go/src/StartloJ/graylog-alert-exporter
COPY .   .
ARG VERSION=dev
ARG COMMIT=dev
ARG DATE=dev
RUN CGO_ENABLED=0 go build -ldflags "-s -w -X main.version=$VERSION -X main.commit=$COMMIT -X main.date=$DATE" -i -a -o app .

FROM gcr.io/distroless/base
WORKDIR /
COPY --from=builder /go/src/StartloJ/graylog-alert-exporter/app /
EXPOSE 9889
ENTRYPOINT ["/app"]