package stderrorhandler

import "log"

func StdErrorHandler(msg string, err error) error {
	log.Println(msg)
	return err
}
