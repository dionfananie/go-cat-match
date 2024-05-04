package helper

import (
	"errors"
)

const (
	MALE   string = "male"
	FEMALE string = "female"
)

func ValidateSex(status string) error {
	switch status {
	case MALE, FEMALE:
		return nil
	default:
		return errors.New("invalid Sex Status, please check your Gender type")
	}
}
