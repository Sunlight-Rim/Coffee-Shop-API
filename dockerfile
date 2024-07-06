# Build stage

FROM golang:alpine3.19 AS builder

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s" -o build/coffeeshop-api cmd/coffeeshop/main.go

# Run stage

FROM alpine:3 AS main

WORKDIR /app

RUN mkdir configs

COPY --from=builder app/build/coffeeshop-api .
COPY --from=builder app/configs/config.json configs/

CMD ["./coffeeshop-api"]