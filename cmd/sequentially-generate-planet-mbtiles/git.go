package sequentiallygenerateplanetmbtiles

import (
	"path/filepath"

	"github.com/lambdajack/sequentially-generate-planet-mbtiles/internal/git"
	"github.com/lambdajack/sequentially-generate-planet-mbtiles/internal/system"
)

type repos struct {
	gdal       git.Repo
	osmiumTool git.Repo
	libosmium  git.Repo
	tilemaker  git.Repo
	tippecanoe git.Repo
}

var gh repos

func cloneRepos() {
	gh = repos{
		gdal: git.Repo{
			Url: "https://github.com/lambdajack/gdal",
			Dst: filepath.Join(pth.temp, "gdal"),
		},
		osmiumTool: git.Repo{
			Url: "https://github.com/lambdajack/osmium-tool",
			Dst: filepath.Join(pth.temp, "osmium-tool"),
		},
		libosmium: git.Repo{
			Url: "https://github.com/lambdajack/libosmium",
			Dst: filepath.Join(pth.temp, "libosmium"),
		},
		tilemaker: git.Repo{
			Url: "https://github.com/lambdajack/tilemaker",
			Dst: filepath.Join(pth.temp, "tilemaker"),
		},
		tippecanoe: git.Repo{
			Url: "https://github.com/lambdajack/tippecanoe",
			Dst: filepath.Join(pth.temp, "tippecanoe"),
		},
	}

	var f []string

	err := gh.gdal.Clone()
	if err != nil {
		f = append(f, "gdal")
	}
	system.SetUserOwner(gh.gdal.Dst)

	err = gh.osmiumTool.Clone()
	if err != nil {
		f = append(f, "osmium-tool")
	}
	system.SetUserOwner(gh.osmiumTool.Dst)

	err = gh.libosmium.Clone()
	if err != nil {
		f = append(f, "libosmium")
	}
	system.SetUserOwner(gh.libosmium.Dst)

	err = gh.tilemaker.Clone()
	if err != nil {
		f = append(f, "tilemaker")
	}
	system.SetUserOwner(gh.tilemaker.Dst)

	err = gh.tippecanoe.Clone()
	if err != nil {
		f = append(f, "tippecanoe")
	}
	system.SetUserOwner(gh.tippecanoe.Dst)

	for _, e := range f {
		lg.err.Fatalf("error cloning %s: %v", e, err)
	}
}
