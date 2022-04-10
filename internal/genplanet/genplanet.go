package genplanet

import (
	"github.com/lambdajack/sequentially-generate-planet-mbtiles/internal/buildthirdpartycontainers"
	"github.com/lambdajack/sequentially-generate-planet-mbtiles/internal/folders"
	"github.com/lambdajack/sequentially-generate-planet-mbtiles/pkg/runtilejoin"
)

func GenPlanet() {
	runtilejoin.RunTileJoin(folders.MbtilesFolder, folders.MbtilesMergedFolder, buildthirdpartycontainers.ContainerTippecanoeName)
}
