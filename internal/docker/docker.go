package docker

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

type Container struct {
	Built      bool
	Context    string
	Dockerfile string
	Flags      []string
	Name       string
	Volumes    []Volume
}

type Volume struct {
	Container string
	Host      string
}

func (c *Container) Build() error {
	if c.Built {
		return nil
	}

	err := exec.Command("docker", "image", "inspect", c.Name).Run()
	if err == nil {
		log.Printf("docker image %s already built - proceeding without cached rebuild", c.Name)
		c.Built = true
		return nil
	}

	if c.Name == "" {
		return fmt.Errorf("container name is empty")
	}

	if c.Dockerfile == "" {
		return fmt.Errorf("dockerfile is empty")
	}

	if c.Context == "" {
		return fmt.Errorf("context is empty")
	}

	cmd := exec.Command("docker", "build", "-t", c.Name, "-f", c.Dockerfile, c.Context)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return err
	}
	cmd.Wait()
	c.Built = true
	return nil
}

func (c *Container) Execute(command []string) error {

	if !c.Built {
		err := c.Build()
		if err != nil {
			return err
		}
	}

	ex := []string{"docker", "run", "--name", c.Name}
	ex = append(ex, c.Flags...)
	for _, v := range c.Volumes {
		ex = append(ex, "-v", v.Host+":"+v.Container)
	}
	ex = append(ex, c.Name)
	ex = append(ex, command...)

	cmd := exec.Command(ex[0], ex[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return err
	}
	cmd.Wait()
	return nil
}

func (c Container) Clean() error {

	out, err := exec.Command("docker", "ps", "-q", "--filter", "name="+c.Name).Output()
	if err != nil {
		return err
	}
	if len(out) < 1 {
		return nil
	}

	log.Printf("attempting to clean up container: %s, ( %s)\n", c.Name, strings.ReplaceAll(string(out), "\n", " "))

	cmd := exec.Command("docker", "stop", c.Name)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {

		return err
	}
	cmd.Wait()
	return nil
}

// Returns a container with the --rm flag set by default.
func New(c Container) *Container {
	rm := false
	for _, f := range c.Flags {
		if f == "--rm" {
			rm = true
		}
	}

	if !rm {
		c.Flags = append(c.Flags, "--rm")
	}

	c.Built = false
	return &c
}
