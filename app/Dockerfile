FROM golang:1.18 AS build

WORKDIR /go/app/cpayment

COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY main.go ./
COPY models models
COPY handlers handlers

RUN go install ./

EXPOSE 9090