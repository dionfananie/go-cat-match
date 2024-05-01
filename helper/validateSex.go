package helper

import (
	"errors"
	"web/go-cat-match/model/cat"
)

const (
	MALE   cat.Sex = "male"
	FEMALE cat.Sex = "female"
)

func ValidateSex(status cat.Sex) error {
	switch status {
	case MALE, FEMALE:
		return nil
	default:
		return errors.New("invalid status")
	}
}
