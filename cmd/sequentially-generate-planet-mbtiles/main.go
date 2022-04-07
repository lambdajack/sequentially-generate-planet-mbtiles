package main

import (
	"github.com/lambdajack/sequentially-generate-planet-mbtiles/internal/buildthirdpartycontainers"
	"github.com/lambdajack/sequentially-generate-planet-mbtiles/internal/clonerepos"
	"github.com/lambdajack/sequentially-generate-planet-mbtiles/internal/downloadosmdata"
	"github.com/lambdajack/sequentially-generate-planet-mbtiles/internal/flags"
	"github.com/lambdajack/sequentially-generate-planet-mbtiles/internal/folders"
	"github.com/lambdajack/sequentially-generate-planet-mbtiles/internal/unzipwaterpolygons"
)

// embed docker files into binary
func main() {
	flags.GetFlags()
	folders.SetupFolderStructure()
	clonerepos.CloneRepos()
	buildthirdpartycontainers.BuildContainers()
	downloadosmdata.DownloadOsmData()
	unzipwaterpolygons.UnzipWaterPolygons()
}
