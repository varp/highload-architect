# syntax=docker/dockerfile:1.2
FROM golang:1.20 AS builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 go build -v -o social-net-server ./cmd/server

FROM debian:bullseye-slim
COPY --from=builder /app/social-net-server /usr/local/bin/social-net-server
EXPOSE 8000
CMD ["social-net-server"]
