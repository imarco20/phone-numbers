FROM golang:1.16-alpine as builder

RUN mkdir /build
WORKDIR /build

COPY . .

RUN export GO111MODULE=on

RUN CGO_ENABLED=1 GOOS=linux go build -a -installsuffix cgo -o app ./cmd/api

FROM alpine:latest

WORKDIR /app
COPY --from=builder /build/app ./app

CMD ./app