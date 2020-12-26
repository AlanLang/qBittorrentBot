# Build go
FROM golang:latest AS serverBuilder
COPY . /qBittorrentBot
WORKDIR /qBittorrentBot
RUN go build -ldflags "-s -w" -o qBittorrentBot

FROM golang:latest
WORKDIR /app
COPY --from=serverBuilder /qBittorrentBot/qBittorrentBot /app/.
ENTRYPOINT ["./qBittorrentBot"]