FROM golang:1.13-stretch as builder

RUN mkdir /app
WORKDIR /app

COPY go.mod .
RUN go mod download

COPY . /app

RUN go run github.com/prisma/prisma-client-go generate --schema=/app/scripts/prisma/schema.prisma

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o /go/bin/stop-analyzing-api .

FROM ubuntu:18.04

WORKDIR /app

RUN apt-get update && apt-get install -y ca-certificates

COPY --from=builder /go/bin/stop-analyzing-api /app/stop-analyzing-api

CMD ["./stop-analyzing-api", "serve"]