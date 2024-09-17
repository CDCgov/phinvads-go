package app

import (
	"database/sql"
	"fmt"
	"regexp"
	"time"

	codes_go_proto "github.com/google/fhir/go/proto/google/fhir/proto/r5/core/codes_go_proto"
	datatypespb "github.com/google/fhir/go/proto/google/fhir/proto/r5/core/datatypes_go_proto"
	r5pb "github.com/google/fhir/go/proto/google/fhir/proto/r5/core/resources/code_system_go_proto"

	"github.com/CDCgov/phinvads-fhir/internal/database/models/xo"
	"github.com/CDCgov/phinvads-fhir/internal/errors"
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

func newString(value string) *datatypespb.String {
	if value == "" {
		return nil
	}
	return &datatypespb.String{Value: value}
}

func newBoolean(value bool) *datatypespb.Boolean {
	return &datatypespb.Boolean{Value: value}
}

func newDateTime(t time.Time) *datatypespb.DateTime {
	if t.IsZero() {
		return nil
	}
	return &datatypespb.DateTime{
		ValueUs:   t.UnixMicro(),
		Precision: datatypespb.DateTime_MICROSECOND,
	}
}

func newNullableString(ns sql.NullString) *datatypespb.String {
	if ns.Valid {
		return newString(ns.String)
	}
	return nil
}

func newId(value string) *datatypespb.Id {
	return &datatypespb.Id{Value: value}
}

func newMarkdown(value string) *datatypespb.Markdown {
	return &datatypespb.Markdown{Value: value}
}

func newNullableMarkdown(ns sql.NullString) *datatypespb.Markdown {
	if ns.Valid {
		return newMarkdown(ns.String)
	}
	return nil
}

func newUri(value string) *datatypespb.Uri {
	return &datatypespb.Uri{Value: value}
}

func newNullableUri(ns sql.NullString) *datatypespb.Uri {
	if ns.Valid {
		return newUri(ns.String)
	}
	return nil
}

func newXhtml(value string) *datatypespb.Xhtml {
	return &datatypespb.Xhtml{Value: value}
}

func serializeCodeSystemToFhir(cs *xo.CodeSystem) (*r5pb.CodeSystem, error) {
	fhirCS := &r5pb.CodeSystem{
		Id:           newId(cs.Oid),
		Status:       &r5pb.CodeSystem_StatusCode{Value: codes_go_proto.PublicationStatusCode_DRAFT},
		Version:      newString(cs.Version),
		Name:         newString(cs.Name),
		Description:  newNullableMarkdown(cs.Definitiontext),
		Experimental: newBoolean(cs.Legacyflag),
		Url:          newNullableUri(cs.Sourceurl),
		Date:         newDateTime(cs.Statusdate),
		Publisher:    newNullableString(cs.Distributionsourceversionname),
		Title:        newNullableString(cs.Assigningauthorityversionname),
	}

	fhirCS.Identifier = []*datatypespb.Identifier{
		{
			System: newUri("urn:ietf:rfc:3986"), // Assuming this system for oid mapping
			Value:  newString(fmt.Sprintf("urn:oid:%s", cs.Oid)),
		},
	}

	fhirCS.Meta = &datatypespb.Meta{
		Profile: []*datatypespb.Canonical{
			{Value: "http://hl7.org/fhir/StructureDefinition/shareablecodesystem"},
		},
	}

	fhirCS.Text = &datatypespb.Narrative{
		Status: &datatypespb.Narrative_StatusCode{Value: codes_go_proto.NarrativeStatusCode_GENERATED},
		Div:    newXhtml("<div>Your narrative text here</div>"),
	}

	return fhirCS, nil
}
