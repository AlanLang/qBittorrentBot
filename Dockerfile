# Build go
FROM golang:1.13-alpine as builder
COPY . /qBittorrentBot
RUN apk add git make gcc libc-dev && \
    cd /qBittorrentBot && make build

FROM alpine:latest
WORKDIR /app
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /qBittorrentBot/qBittorrentBot /app/.
VOLUME /app/config
ENTRYPOINT ["./qBittorrentBot"]