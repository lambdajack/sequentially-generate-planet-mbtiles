package main

import (
	"github.com/lambdajack/sequentially-generate-planet-mbtiles/internal/buildthirdpartycontainers"
	"github.com/lambdajack/sequentially-generate-planet-mbtiles/internal/clonerepos"
	"github.com/lambdajack/sequentially-generate-planet-mbtiles/internal/downloadosmdata"
	"github.com/lambdajack/sequentially-generate-planet-mbtiles/internal/extractquadrants"
	"github.com/lambdajack/sequentially-generate-planet-mbtiles/internal/extractslices"
	"github.com/lambdajack/sequentially-generate-planet-mbtiles/internal/flags"
	"github.com/lambdajack/sequentially-generate-planet-mbtiles/internal/folders"
	"github.com/lambdajack/sequentially-generate-planet-mbtiles/internal/genmbtiles"
	"github.com/lambdajack/sequentially-generate-planet-mbtiles/internal/genplanet"
	"github.com/lambdajack/sequentially-generate-planet-mbtiles/internal/unzippolygons"
)

func main() {
	flags.GetFlags()
	folders.SetupFolderStructure()
	clonerepos.CloneRepos()
	buildthirdpartycontainers.BuildContainers()
	downloadosmdata.DownloadOsmData()
	unzippolygons.UnzipPolygons()
	extractquadrants.ExtractQuadrants()
	extractslices.ExtractSlicesFromQuadrants()
	genmbtiles.GenMbtiles()
	genplanet.GenPlanet()
}
