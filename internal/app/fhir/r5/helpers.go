package r5

import (
	"database/sql"
	"time"

	datatypespb "github.com/google/fhir/go/proto/google/fhir/proto/r5/core/datatypes_go_proto"
)

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

//nolint:unused
func newNullableUri(ns sql.NullString) *datatypespb.Uri {
	if ns.Valid {
		return newUri(ns.String)
	}
	return nil
}

func newXhtml(value string) *datatypespb.Xhtml {
	return &datatypespb.Xhtml{Value: value}
}

func newUnsignedInt(value int) *datatypespb.UnsignedInt {
	return &datatypespb.UnsignedInt{Value: uint32(value)}
}
