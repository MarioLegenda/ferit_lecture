FROM golang:1.18-alpine as golang_build

ENV APP_DIR /app
WORKDIR /app

VOLUME $APP_DIR

RUN apk add build-base
RUN apk add --update --no-cache --repository http://dl-3.alpinelinux.org/alpine/edge/community --repository http://dl-3.alpinelinux.org/alpine/edge/main vips-dev

COPY ./go.mod .
COPY ./go.sum .
RUN go mod download && go mod tidy

RUN go install github.com/githubnemo/CompileDaemon@latest

RUN go env

COPY ./ .

EXPOSE 3002

ENTRYPOINT /go/bin/CompileDaemon --build="go build -o go_build" --command=./go_build --directory="/app" --pattern="(.+\.go)$" --polling --polling-interval=1000 --graceful-kill=true --color=true