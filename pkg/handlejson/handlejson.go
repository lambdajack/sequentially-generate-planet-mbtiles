package handlejson

import (
	"encoding/json"
	"errors"
	"io"
	"os"
)

func DecodeTo(j interface{}, jsonPath string) error {
	jr, err := os.Open(jsonPath)
	if err != nil {
		return err
	}

	dec := json.NewDecoder(jr)
	dec.DisallowUnknownFields()

	err = dec.Decode(&j)
	if err != nil {
		return err
	}

	err = dec.Decode(&struct{}{})
	if err != io.EOF {
		return errors.New("supplied more than one json object when only one is allowed")
	}

	return nil
}
