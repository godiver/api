FROM golang:1.19.0-alpine3.16 as build

ARG RAKUTEN_APP_ID=""
ARG YOUTUBE_DEVELOPER_KEY=""

ENV GO111MODULE="on" \
    APP="/go/src/github.com/godiver/api" \
    RAKUTEN_APP_ID=$RAKUTEN_APP_ID \
    YOUTUBE_DEVELOPER_KEY=$YOUTUBE_DEVELOPER_KEY

WORKDIR ${APP}

RUN apk update && \
    apk upgrade && \
    apk add --no-cache \
    git && \
    go install github.com/pilu/fresh@latest

COPY ./ .

RUN go build -o app ${APP}/main.go

# 本番用のDocker imageを作成する
FROM alpine:3.12.1
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=build /go/src/github.com/godiver/api/app .
EXPOSE 8080
CMD ["./app"]
