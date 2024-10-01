package models

import (
	"context"

	"github.com/CDCgov/phinvads-go/internal/database/models/xo"
)

func GetAllCodeSystems(ctx context.Context, db xo.DB) (*[]xo.CodeSystem, error) {
	const sqlstr = `SELECT * FROM public.code_system`

	codeSystems := []xo.CodeSystem{}
	rows, err := db.QueryContext(ctx, sqlstr)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		cs := xo.CodeSystem{}
		err := rows.Scan(&cs.Oid, &cs.ID, &cs.Name, &cs.Definitiontext, &cs.Status, &cs.Version, &cs.Versiondescription, &cs.Assigningauthorityversionname, &cs.Distributionsourceversionname, &cs.Distributionsourceid, &cs.Assigningauthorityid, &cs.Codesystemcode, &cs.Sourceurl, &cs.Hl70396identifier, &cs.Legacyflag, &cs.Statusdate, &cs.Acquireddate, &cs.Effectivedate, &cs.Expirydate, &cs.Assigningauthorityreleasedate, &cs.Distributionsourcereleasedate, &cs.Sdocreatedate, &cs.Lastrevisiondate)
		if err != nil {
			return nil, err
		}
		codeSystems = append(codeSystems, cs)
	}
	return &codeSystems, nil
}

func GetCodeSystemByLikeOID(ctx context.Context, db xo.DB, oid string) (*[]xo.CodeSystem, error) {
	wildcard := "%" + oid + "%"
	const sqlstr = "SELECT * FROM public.code_system WHERE oid LIKE $1"

	return queryCodeSystems(ctx, db, sqlstr, wildcard)
}

func GetCodeSystemByLikeID(ctx context.Context, db xo.DB, id string) (*[]xo.CodeSystem, error) {
	wildcard := "%" + id + "%"
	const sqlstr = "SELECT * FROM public.code_system WHERE id LIKE $1"

	return queryCodeSystems(ctx, db, sqlstr, wildcard)
}

func GetCodeSystemByLikeString(ctx context.Context, db xo.DB, input string) (*[]xo.CodeSystem, error) {
	wildcard := "%" + input + "%"
	const sqlstr = "SELECT * FROM public.code_system WHERE lower(name) LIKE lower($1) OR lower(codesystemcode) LIKE lower ($1)"

	return queryCodeSystems(ctx, db, sqlstr, wildcard)
}

func queryCodeSystems(ctx context.Context, db xo.DB, sqlstr, wildcard string) (*[]xo.CodeSystem, error) {
	codeSystems := []xo.CodeSystem{}
	rows, err := db.QueryContext(ctx, sqlstr, wildcard)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		cs := xo.CodeSystem{}
		err := rows.Scan(&cs.Oid, &cs.ID, &cs.Name, &cs.Definitiontext, &cs.Status, &cs.Version, &cs.Versiondescription, &cs.Assigningauthorityversionname, &cs.Distributionsourceversionname, &cs.Distributionsourceid, &cs.Assigningauthorityid, &cs.Codesystemcode, &cs.Sourceurl, &cs.Hl70396identifier, &cs.Legacyflag, &cs.Statusdate, &cs.Acquireddate, &cs.Effectivedate, &cs.Expirydate, &cs.Assigningauthorityreleasedate, &cs.Distributionsourcereleasedate, &cs.Sdocreatedate, &cs.Lastrevisiondate)
		if err != nil {
			return nil, err
		}
		codeSystems = append(codeSystems, cs)
	}
	return &codeSystems, nil
}

type SearchResultRow struct {
	CodeSystemsCount        string
	CodeSystemConceptsCount string
	ValueSetsCount          string
	ValueSetConceptsCount   string
	CodeSystems             []*xo.CodeSystem
	CodeSystemConcepts      []*xo.CodeSystemConcept
	ValueSets               []*xo.ValueSet
	ValueSetConcepts        []*xo.ValueSet
	URL                     string
	PageCount               int
}
