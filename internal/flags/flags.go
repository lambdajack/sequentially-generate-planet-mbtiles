package flags

import (
	"flag"
)

func GetFlags() (pathToConfig *string) {
	pathToConfig = flag.String("c", "config.json", "Relative or absolute path to config.json file")
	flag.Parse()
	return pathToConfig
}
