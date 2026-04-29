FROM golang:1.26.2-alpine AS build
WORKDIR /app

COPY go.sum go.mod  ./
RUN go mod download

