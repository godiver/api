FROM golang:1.14.6-alpine3.12 as build

ARG RAKUTEN_APP_ID=""

ENV GO111MODULE="on" \
    APP="/go/src/github.com/godiver/api" \
    RAKUTEN_APP_ID=$RAKUTEN_APP_ID

WORKDIR ${APP}

RUN apk update && \
    apk upgrade && \
    apk add --no-cache \
    git && \
    go get github.com/pilu/fresh

COPY ./ .

RUN go build -o app ${APP}/main.go
