package system

import (
	"log"
	"os"
	"os/exec"
	"os/user"
	"runtime"
	"strconv"
	"strings"
)

func SetUserOwner(path string) error {
	u, err := user.Current()
	if err != nil {
		return err
	}

	if u.Name == "root" && runtime.GOOS == "linux" {
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

func UserHomeDir() string {
	if runtime.GOOS == "linux" {
		if u, err := user.Current(); err == nil {
			if u.Username != "root" {
				return u.HomeDir
			}
			username := os.Getenv("SUDO_USER")
			if username != "" {
				u, err := user.Lookup(username)
				if err != nil {
					return ""
				}
				return u.HomeDir
			}
		}
	}
	d, err := os.UserHomeDir()
	if err != nil {
		return ""
	}
	return d
}

func UserCacheDir() string {
	if runtime.GOOS == "linux" {
		if UserHomeDir() != "" {
			return UserHomeDir() + "/.cache"
		}
		return ""
	}
	d, err := os.UserCacheDir()
	if err != nil {
		return ""
	}
	return d
}

func DockerIsSnap() bool {
	if runtime.GOOS == "linux" {
		sl, err := exec.Command("snap", "list", "docker").CombinedOutput()
		if err != nil {
			return false
		}
		return strings.Contains(string(sl), "docker")
	}
	return false
}
