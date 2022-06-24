VERSION := v3.1.0-rc10

all: darwin linux windows

darwin:
	GOOS=darwin go build -o bin/sequentially-generate-planet-mbtiles--darwin-$(VERSION) -ldflags "-X github.com/lambdajack/sequentially-generate-planet-mbtiles/cmd/sequentially-generate-planet-mbtiles.sgpmVersion=$(VERSION)"

linux:
	GOOS=linux go build -o bin/sequentially-generate-planet-mbtiles--unix-$(VERSION) -ldflags "-X github.com/lambdajack/sequentially-generate-planet-mbtiles/cmd/sequentially-generate-planet-mbtiles.sgpmVersion=$(VERSION)"

windows:
	GOOS=windows go build -o bin/sequentially-generate-planet-mbtiles--win-$(VERSION).exe -ldflags "-X github.com/lambdajack/sequentially-generate-planet-mbtiles/cmd/sequentially-generate-planet-mbtiles.sgpmVersion=$(VERSION)"

