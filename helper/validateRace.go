package helper

import (
	"errors"
)

const (
	PERSIAN    string = "Persian"
	MAINE_COON string = "Maine Coon"
	SIAMESE    string = "Siamese"
	RAGDOLL    string = "Ragdoll"
	BENGAL     string = "Bengal"
	SPHYNX     string = "Sphynx"
	BRITISH    string = "British Shorthair"
	ABYSSINIAN string = "Abyssinian"
	SCOTTISH   string = "Scottish Fold"
	BIRMAN     string = "Birman"
)

func ValidateRace(status string) error {
	switch status {
	case PERSIAN, MAINE_COON,
		SIAMESE,
		RAGDOLL,
		BENGAL,
		SPHYNX,
		BRITISH,
		ABYSSINIAN,
		SCOTTISH,
		BIRMAN:
		return nil
	default:
		return errors.New("invalid Race Type, please check your Cat's Race")
	}
}
