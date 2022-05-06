package execute

// func OutputToConsole(commandString string) error {
// 	comArgs := strings.Split(strings.Trim(commandString, " "), " ")
// 	if len(comArgs) < 2 {
// 		return stderrorhandler.StdErrorHandler(fmt.Sprintf("execute.go | Received a command, but no arguments... exiting. \n Received the following command: %v", commandString), nil)
// 	}

// 	cmd := exec.Command(comArgs[0], comArgs[1:]...)
// 	cmd.Stdout = os.Stdout
// 	cmd.Stderr = os.Stderr

// 	if err := cmd.Run(); err != nil {
// 		return stderrorhandler.StdErrorHandler(fmt.Sprintf("execute.go | Executing %v failed.", cmd), err)
// 	}
// 	cmd.Wait()
// 	return nil
// }
