FROM golang:1.20-alpine3.17 as builder

RUN apk update && \
    apk add --no-cache file build-base

WORKDIR /go/src/app

COPY .. .
run mkdir /app/

RUN go mod download
RUN CGO_ENABLED=0 go build -o /app/telegram_bot ./init/main.go

# deploy-stage
FROM alpine:latest

RUN apk update && \
    apk add --no-cache curl bash

RUN mkdir  /key

WORKDIR /app
VOLUME /data

COPY --from=builder /app ./

RUN apk update && \
    apk add --no-cache curl bash

ENTRYPOINT ["/app/telegram_bot"]
