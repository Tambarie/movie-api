FROM golang:1.17.0-alpine

WORKDIR /app

COPY ./ /app

RUN go get github.com/go-redis/redis

RUN go mod tidy

ENTRYPOINT [ "go", "run", "movie-service" ]