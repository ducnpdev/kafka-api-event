FROM golang:1.21.1-alpine AS builder
WORKDIR /app
COPY . /app

ENV BUILD_TAG 1.0.0
ENV GO111MODULE on
ENV CGO_ENABLED=0
ENV GOOS=linux
RUN go mod vendor
RUN go build -o event-tracking /app/cmd/server/main.go

# stage2.1: rebuild
FROM alpine
WORKDIR /app
RUN ls
COPY --from=builder /app/event-tracking /app/event-tracking.go
# COPY --from=builder /app/.env /app/.env
COPY --from=builder /app/config/dev.yaml /app/config/dev.yaml
CMD ["./event-tracking.go"]