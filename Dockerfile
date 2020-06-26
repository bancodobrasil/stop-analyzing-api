FROM golang:1.12.5-stretch as builder

RUN mkdir /app
WORKDIR /app

COPY ./src/go.mod ./src/go.sum ./src/
RUN cd /app/src \ 
    && go mod download

COPY ./src /app/src
COPY ./prisma /app/prisma

RUN cd /app/src \
    && go run github.com/prisma/prisma-client-go generate --schema=/app/prisma/schema.prisma

RUN cd /app/src \
    # && go test ./service \
    && CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o /go/bin/stop-analyzing-api .

FROM ubuntu:18.04

WORKDIR /app

RUN apt-get update && apt-get install -y ca-certificates

COPY --from=builder /go/bin/stop-analyzing-api /app/stop-analyzing-api

CMD ["./stop-analyzing-api", "serve"]