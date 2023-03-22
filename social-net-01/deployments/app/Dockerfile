# syntax=docker/dockerfile:1.2
FROM golang:1.19 as builder
WORKDIR /app
COPY . .
RUN go build -o social-net-server ./cmd/server

FROM alpine:latest
COPY --from=builder /app/social-net-server /usr/local/bin/
ENTRYPOINT ["social-net-server"]
