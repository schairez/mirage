FROM golang:1.15.2-alpine3.12 AS builder

LABEL maintainer="Sergio Chairez <schairezv@gmail.com>"

# ARG BUILD_APP_VERSION="dev version on docker"
# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /app

COPY go.mod go.sum ./

RUN apk add --no-cache git && \
    go mod download

COPY . .
ENV PORT 5000

# RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build
# RUN go run cmd/app/main.go 

CMD [ "go", "run", "cmd/app/main.go" ]

# CMD [ "./mirage" ]
