FROM golang:1.20-alpine

WORKDIR /usr/src/app

ARG env

COPY go.mod go.sum ./

RUN apk update
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /usr/local/bin/app ./...

CMD ["app"]
