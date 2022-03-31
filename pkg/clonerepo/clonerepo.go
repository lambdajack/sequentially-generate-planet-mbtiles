package clonerepo

import (
	"log"
	"os"
	"os/exec"
	"strings"
)

func CloneRepo(repoToClone string) error {
	// Ensure a string was provided
	if repoToClone == "" {
		panic("lambdajack/pkg/clonerepo | Received an empty string (no repo to clone)\n")
	}

	// Check to see if the repo dir already exists, and skip if so.
	splitPath := strings.Split(repoToClone, "/")
	if _, err := os.Stat(splitPath[len(splitPath)-1]); !os.IsNotExist(err) {
		log.Printf("%v already exists. Skipping clone.\n", splitPath[len(splitPath)-1])
		return nil
	}

	// Attempt to validate the url to clone the repo from and correct it if necessary.
	if strings.HasPrefix(repoToClone, "github.com/") {
		repoToClone = "https://" + repoToClone
	}
	if !strings.HasPrefix(repoToClone, "https://github.com/") && !strings.HasSuffix(repoToClone, "http://github.com/") {
		repoToClone = "https://github.com/" + repoToClone
	}

	// Execute git clone repoUrl
	cmd := exec.Command("git", "clone", repoToClone)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Start(); err != nil {
		return err
	}
	cmd.Wait()
	return nil
}
