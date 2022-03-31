package execute

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func OutputToConsole(commandString string) error {
	comArgs := strings.Split(strings.Trim(commandString, " "), " ")
	if len(comArgs) < 2 {
		errStr := fmt.Errorf("pkg/execute | Received a command, but no arguments... exiting. \n Received the following command: %v", commandString)
		return errStr
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
