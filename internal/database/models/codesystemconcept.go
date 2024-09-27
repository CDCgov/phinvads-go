package models

import (
	"context"

	"github.com/CDCgov/phinvads-go/internal/database/models/xo"
)

// All retrieves all rows from 'public.code_system_concept'
func GetAllCodeSystemConcepts(ctx context.Context, db xo.DB) (*[]xo.CodeSystemConcept, error) {
	const sqlstr = `SELECT * FROM public.code_system_concept`
	codeSystemConcepts := []xo.CodeSystemConcept{}
	rows, err := db.QueryContext(ctx, sqlstr)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		csc := xo.CodeSystemConcept{}
		err := rows.Scan(&csc.ID, &csc.Name, &csc.Codesystemoid, &csc.Conceptcode, &csc.Definitiontext, &csc.Status, &csc.Statusdate)
		if err != nil {
			return nil, err
		}
		codeSystemConcepts = append(codeSystemConcepts, csc)
	}
	return &codeSystemConcepts, nil

}
