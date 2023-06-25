FROM golang:1.20-alpine

WORKDIR /usr/src/app

ARG env

COPY go.mod go.sum ./

RUN apk update
RUN apk add postgresql
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /usr/local/bin/app ./main.go

CMD ["app"]
