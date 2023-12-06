# syntax=docker/dockerfile:1

FROM golang:1.19-alpine

WORKDIR /backend

COPY . .

RUN go mod download

RUN go build -o /main

EXPOSE 3000

CMD [ "/main" ]