package git

import (
	"log"
	"os"
	"os/exec"
)

type Repo struct {
	Url string
	Dst string
}

func (r Repo) Clone() error {
	if _, err := os.Stat(r.Dst); err == nil {
		log.Printf("git repo %s already exists, skipping clone", r.Dst)
		return nil
	}

	cmd := exec.Command("git", "clone", r.Url, r.Dst)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return err
	}
	cmd.Wait()

	return nil
}
