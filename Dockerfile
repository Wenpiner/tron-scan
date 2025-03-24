FROM golang:1.21-alpine AS builder


WORKDIR /app

COPY . ./
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o tron-scan .

FROM alpine:3.18.4

WORKDIR /app

ENV TZ=UTC
RUN apk update --no-cache && apk add --no-cache tzdata

COPY --from=builder /app/tron-scan ./
COPY ./etc/tron-api.yaml ./etc/

ENTRYPOINT ["/app/tron-scan", "-f", "/app/etc/tron-api.yaml"]
