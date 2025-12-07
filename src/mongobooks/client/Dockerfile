ARG BUILDER_IMAGE=golang:1.22.5-alpine3.20
FROM ${BUILDER_IMAGE} AS builder

WORKDIR /app
COPY ./ ./
RUN go mod download
RUN go build -o mongobooks ./cmd/api/main.go

FROM alpine:3.23

ENV TZ Europe/Moscow
WORKDIR /app

COPY --from=builder /app/mongobooks /app/mongobooks 

EXPOSE 8080
CMD ["./mongobooks"]
