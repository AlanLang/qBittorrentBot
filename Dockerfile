# Build go
FROM golang:1.15 AS serverBuilder
COPY . /qBittorrentBot
WORKDIR /qBittorrentBot
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o qBittorrentBot

FROM scratch

COPY --from=serverBuilder /qBittorrentBot/qBittorrentBot .

# 这里跟编译完的文件名一致
ENTRYPOINT  ["./qBittorrentBot"]