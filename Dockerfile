# syntax=docker/dockerfile:1
FROM golang:1.17-alpine

WORKDIR /server

COPY go.mod ./
COPY go.sum ./

RUN go mod download
COPY ./ ./

RUN go build -o /lost-n-found

EXPOSE 3000

CMD [ "/lost-n-found" ]