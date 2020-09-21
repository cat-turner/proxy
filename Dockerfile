FROM golang:1.14-alpine
ENV GO111MODULE=on
RUN apk update && apk upgrade && \
    apk add --no-cache bash

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o ./proxys .

RUN chmod a+x ./proxys
EXPOSE $PORT
ENTRYPOINT ["./proxys"]