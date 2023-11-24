## Build
FROM golang:1.21-alpine AS build

WORKDIR /app

RUN apk --no-cache update && \
apk --no-cache add git gcc libc-dev

COPY go.mod go.sum  ./
RUN go mod download

COPY . .

# Kafka Go client is based on the C library librdkafka
ENV CGO_ENABLED 1
ENV GOOS=linux
ENV GOARCH=amd64

RUN export GO111MODULE=on

RUN go build -tags musl

## Dev
FROM build AS dev

WORKDIR /app

RUN apk add --no-cache make

RUN go install github.com/cespare/reflex@latest

ENTRYPOINT ["./entry.sh"]
CMD ["make watch"]


## Prod
FROM alpine:latest AS prod

WORKDIR /

COPY --from=build /app/server /app/entry.sh /app/.env  /

ENTRYPOINT ["/entry.sh"]
CMD ["./server"]