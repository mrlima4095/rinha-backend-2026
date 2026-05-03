FROM golang:1.26.2-alpine AS autogen
WORKDIR /app

COPY  go.mod go.sum  ./
RUN go mod download

# Tools for generation
COPY tools ./tools
COPY internal/references internal/references
COPY pkg/vector pkg/vector

RUN mkdir -p internal/references
RUN go generate ./...

FROM golang:1.26.2-alpine AS build
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -ldflags="-s -w" -trimpath -buildvcs=false -o app ./cmd/app

FROM alpine:3.20 AS certs
RUN apk --no-cache add ca-certificates

FROM scratch AS runner
WORKDIR /srv
COPY --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

COPY --from=autogen /app/references_database .

COPY --from=build /app/app . 


ENTRYPOINT [ "./app" ]

EXPOSE 8080