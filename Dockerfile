FROM golang:1.24.2-alpine3.21 AS build

RUN apk --no-cache add curl

WORKDIR /go/src/web

COPY .. .

RUN go mod download

ENV GO111MODULE=on

RUN go build main.go

FROM alpine:3.21
COPY --from=build /go/src/web/main /app/main
WORKDIR /app

ENTRYPOINT ["./main"]