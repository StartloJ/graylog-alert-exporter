FROM golang:1.16.3-alpine as builder
WORKDIR /go/src/StartloJ/graylog-alert-exporter
COPY .   .
ARG VERSION=dev
ARG COMMIT=dev
ARG DATE=dev
RUN CGO_ENABLED=0 go build -ldflags "-s -w -X main.GitCommit=$COMMIT" -i -a -o app .
CMD ["./git-tester"]

FROM gcr.io/distroless/base
WORKDIR /
COPY --from=builder /go/src/StartloJ/graylog-alert-exporter/app /
EXPOSE 9889
ENTRYPOINT ["/app"]