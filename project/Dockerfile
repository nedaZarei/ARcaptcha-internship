FROM golang:1.23.2-alpine AS builder

RUN apk add --no-cache git build-base

WORKDIR /app
COPY go.mod go.sum ./
COPY ./config/config.yaml ./
ENV GOPROXY=https://goproxy.cn
RUN go mod download
COPY . .
RUN go build -o apartment

FROM builder AS tester
RUN go install gotest.tools/gotestsum@latest
RUN go install github.com/vektra/mockery/v2@latest

FROM alpine:latest
RUN apk add --no-cache curl
COPY --from=builder /app/apartment /app/apartment
COPY --from=builder /app/config.yaml /app/config.yaml
EXPOSE 8080
WORKDIR /app
CMD ["/app/apartment", "start"]