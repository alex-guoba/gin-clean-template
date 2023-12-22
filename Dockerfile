# Build stage
FROM golang:1.21-alpine3.18 AS builder
WORKDIR /app
COPY . .
RUN go build -o gin-clean-template main.go

# RUN State
FROM alpine:3.19
WORKDIR /app
COPY --from=builder /app/gin-clean-template .
COPY configs ./configs
COPY db/migration ./db/migration
COPY scripts ./scripts

RUN apk add --no-cache tzdata
ENV TZ="UTC"

EXPOSE 8080

ENTRYPOINT ["/app/gin-clean-template"]