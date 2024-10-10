package app

import (
	"regexp"

	"github.com/CDCgov/phinvads-go/internal/errors"
)

type IdType int

const (
	Unknown IdType = iota
	Oid     IdType = iota
	Id      IdType = iota
)

func determineIdType(input string) (output IdType, err error) {
	validId, _ := regexp.MatchString("^[a-zA-Z0-9-]+$", input)
	validOid, _ := regexp.MatchString("^[0-9.]+$", input)

	switch {
	case validId:
		return Id, nil
	case validOid:
		return Oid, nil
	default:
		return Unknown, errors.ErrInvalidId
	}
}
