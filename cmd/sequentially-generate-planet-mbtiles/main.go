package main

import (
	"github.com/lambdajack/sequentially-generate-planet-mbtiles/internal/buildthirdpartycontainers"
	"github.com/lambdajack/sequentially-generate-planet-mbtiles/internal/downloadosmdata"
	"github.com/lambdajack/sequentially-generate-planet-mbtiles/internal/flags"
	"github.com/lambdajack/sequentially-generate-planet-mbtiles/internal/folders"
)

func main() {
	flags.GetFlags()
	folders.SetupFolderStructure()
	buildthirdpartycontainers.BuildContainers()
	downloadosmdata.DownloadOsmData()
	// unzippolygons.UnzipPolygons()
	// extractquadrants.ExtractQuadrants()
	// extractslices.FromQuadrants()
	// genmbtiles.GenMbtiles()
	// genplanet.GenPlanet()
}
