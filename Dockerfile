FROM golang:1.18

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY main.go ./
COPY handlers models
COPY handlers handlers

RUN go install ./

EXPOSE 5300