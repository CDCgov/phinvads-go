package app

import (
	"regexp"

	"github.com/CDCgov/phinvads-go/internal/errors"
)

func determineParamType(input string) (output string, err error) {
	validUuid, _ := regexp.MatchString("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$", input)
	validOid, _ := regexp.MatchString("^[0-9.]+$", input)
	validString, _ := regexp.MatchString("^[a-zA-Z0-9 ]*$", input)

	if validUuid {
		return "uuid", nil
	} else if validOid {
		return "oid", nil
	} else if validString {
		return "string", nil
	} else {
		return "", errors.ErrInvalidId
	}
}
