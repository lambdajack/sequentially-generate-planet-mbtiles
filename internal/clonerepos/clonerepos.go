package clonerepos

import (
	"fmt"

	"github.com/lambdajack/sequentially-generate-planet-mbtiles/pkg/clonerepo"
	"github.com/lambdajack/sequentially-generate-planet-mbtiles/pkg/stderrorhandler"
)

func CloneRepos() error {
	reposToClone := [...]string{"systemed/tilemaker", "mapbox/tippecanoe"}

	for _, repo := range reposToClone {
		err := clonerepo.CloneRepo(repo)
		if err != nil {
			stderrorhandler.StdErrorHandler(fmt.Sprintf("clonerepos.go | Failed to clone %s. Unable to proceed...", repo), err)
			panic(err)
		}
	}

	return nil
}
