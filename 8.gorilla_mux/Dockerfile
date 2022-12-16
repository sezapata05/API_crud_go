#build stage
FROM golang:1.19.3-alpine3.17 AS builder
RUN apk add --no-cache git

WORKDIR /app

COPY ["go.mod", "go.sum", "./"]
RUN go mod download

COPY . .

RUN go get -d -v ./...

RUN go build -o app -v .

#final stage
FROM alpine:3.17
LABEL Name=gorilla-api Version=0.0.1
RUN apk update
RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY --from=builder /app /app
ENTRYPOINT /app
