package execute

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func OutputToConsole(command string) error {
	comArgs := strings.Split(strings.Trim(command, " "), " ")
	if len(comArgs) < 2 {
		return fmt.Errorf("execute.go | OutputToConsole | Invalid command string: %s", command)
	}

	cmd := exec.Command(comArgs[0], comArgs[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return err
	}
	cmd.Wait()
	return nil
}
