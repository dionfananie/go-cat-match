package helper

import (
	"errors"
	"web/go-cat-match/model/cat"
)

const (
	PERSIAN    cat.Race = "Persian"
	MAINE_COON cat.Race = "Maine Coon"
	SIAMESE    cat.Race = "Siamese"
	RAGDOLL    cat.Race = "Ragdoll"
	BENGAL     cat.Race = "Bengal"
	SPHYNX     cat.Race = "Sphynx"
	BRITISH    cat.Race = "British Shorthair"
	ABYSSINIAN cat.Race = "Abyssinian"
	SCOTTISH   cat.Race = "Scottish Fold"
	BIRMAN     cat.Race = "Birman"
)

func ValidateRace(status cat.Race) error {
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
		return errors.New("invalid status")
	}
}
