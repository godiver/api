FROM golang:1.14.6-alpine3.12 as build

ENV APP="/go/src/github.com/godiver/api"
WORKDIR ${APP}
RUN apk update && \
    apk upgrade && \
    apk add --no-cache \
    git && \
    go get github.com/pilu/fresh
