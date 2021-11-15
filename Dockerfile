# syntax=docker/dockerfile:1

FROM golang:1.17-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./
COPY assets ./assets

RUN go build -o /server

EXPOSE 8080

CMD ["/server"]

