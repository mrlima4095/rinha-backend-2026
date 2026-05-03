VERSION 0.8
FROM golang:1.26.1-alpine
WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY tools ./tools
RUN mkdir -p internal/references
RUN go generate ./...


linux: 
    COPY . .
    RUN GOOS=linux go build -ldflags="-s -w" -trimpath -buildvcs=false ./cmd/app 
    SAVE ARTIFACT app AS LOCAL dist/

tests:
    COPY . . 
    RUN go test -v ./...

