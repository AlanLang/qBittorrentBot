build: get
	go build -ldflags "-X 'github.com/AlanLang/qBittorrentBot/config.commit=`git rev-parse --short HEAD`' -X 'github.com/AlanLang/qBittorrentBot/config.date=`date`'"

get:
	go mod download

run:
	go run .