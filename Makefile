VERSION := v3.1.0-rc3

all: darwin linux windows

darwin:
	GOOS=darwin go build -o bin/sequentially-generate-planet-mbtiles--darwin-$(VERSION)

linux:
	GOOS=linux go build -o bin/sequentially-generate-planet-mbtiles--unix-$(VERSION)

windows:
	GOOS=windows go build -o bin/sequentially-generate-planet-mbtiles--win-$(VERSION).exe

