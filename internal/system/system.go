package system

import (
	"log"
	"os"
	"os/user"
	"strconv"
)

func SetUserOwner(path string) error {
	u, err := user.Current()
	if err != nil {
		return err
	}

	if u.Name == "root" {

		u := os.Getenv("SUDO_UID")
		if u == "" {
			return err
		}
		ui, err := strconv.Atoi(u)
		if err != nil {
			return err
		}

		g := os.Getenv("SUDO_GID")
		if g == "" {
			return err
		}
		gi, err := strconv.Atoi(g)
		if err != nil {
			return err
		}

		err = os.Chown(path, ui, gi)
		if err != nil {
			log.Println("failed to set permissions for directories", path, err)
		}

	}
	return nil
}
