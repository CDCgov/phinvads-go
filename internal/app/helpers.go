package app

import (
	"regexp"

	"github.com/CDCgov/phinvads-go/internal/errors"
)

func determineIdType(input string) (output string, err error) {
	validId, _ := regexp.MatchString("^[a-zA-Z0-9-]+$", input)
	validOid, _ := regexp.MatchString("^[0-9.]+$", input)

	if validId {
		return "id", nil
	} else if validOid {
		return "oid", nil
	} else {
		return "", errors.ErrInvalidId
	}
}
