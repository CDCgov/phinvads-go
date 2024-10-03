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
	wildcard := oid + "%"
	const sqlstr = "SELECT * FROM public.code_system WHERE oid LIKE $1"

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

type CodeSystemResultRow struct {
	CodeSystemsCount        string
	CodeSystemConceptsCount string
	ValueSetsCount          string
	CodeSystems             []*xo.CodeSystem
	CodeSystemConcepts      []*xo.CodeSystemConcept
	ValueSets               []*xo.ValueSet
	URL                     string
	PageCount               int
}

func GetCodeSystemConceptCount(ctx context.Context, db xo.DB, oid string) (int, error) {
	const sqlstr = `SELECT COUNT(*) FROM public.code_system_concept WHERE codesystemoid = $1`

	var count int
	err := db.QueryRowContext(ctx, sqlstr, oid).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}
