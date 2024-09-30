package app

import (
	"regexp"

	"github.com/CDCgov/phinvads-go/internal/errors"
)

func determineParamType(input string) (output string, err error) {
	validId, _ := regexp.MatchString("^[a-zA-Z]+[0-9-]+$", input)
	validOid, _ := regexp.MatchString("^[0-9.]+$", input)
	validString, _ := regexp.MatchString("^[a-zA-Z0-9 ]*$", input)

	if validId {
		return "id", nil
	} else if validOid {
		return "oid", nil
	} else if validString {
		return "string", nil
	} else {
		return "", errors.ErrInvalidId
	}
}
