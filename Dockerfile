FROM golang:1.12.5-stretch as builder

RUN mkdir /app
WORKDIR /app

COPY ./src/go.mod ./src/go.sum ./
RUN go mod download

COPY ./src .

RUN go test ./... \
    && CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o /go/bin/stop-analyzing-api .

FROM alpine:3.12

WORKDIR /app

RUN apk add --quiet --no-cache openssl=1.1.1g-r0

COPY --from=builder /go/bin/stop-analyzing-api /app/stop-analyzing-api

CMD [ "./stop-analyzing-api", "serve" ]