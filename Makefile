VERSION := v3.1.0
LDF := "-X github.com/lambdajack/sequentially-generate-planet-mbtiles/cmd/sequentially-generate-planet-mbtiles.sgpmVersion=$(VERSION)"

all: clean darwin linux windows

darwin:
	GOOS=darwin GOARCH=amd64 go build -o bin/sequentially-generate-planet-mbtiles--darwin-amd64-$(VERSION) -ldflags $(LDF)

linux:
	GOOS=linux GOARCH=amd64 go build -o bin/sequentially-generate-planet-mbtiles--unix-amd64-$(VERSION) -ldflags $(LDF)
	GOOS=linux GOARCH=arm go build -o bin/sequentially-generate-planet-mbtiles--unix-arm64-$(VERSION) -ldflags $(LDF)

windows:
	GOOS=windows GOARCH=amd64 go build -o bin/sequentially-generate-planet-mbtiles--win-amd64-$(VERSION).exe -ldflags $(LDF)

clean:
	rm -rf bin/