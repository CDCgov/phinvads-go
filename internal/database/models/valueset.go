package models

import (
	"context"

	"github.com/CDCgov/phinvads-go/internal/database/models/xo"
)

// All retrieves all rows from 'public.value_set'
func GetAllValueSets(ctx context.Context, db xo.DB) (*[]xo.ValueSet, error) {
	const sqlstr = `SELECT * FROM public.value_set`
	valueSets := []xo.ValueSet{}
	rows, err := db.QueryContext(ctx, sqlstr)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		vs := xo.ValueSet{}
		err := rows.Scan(&vs.Oid, &vs.ID, &vs.Name, &vs.Code, &vs.Status, &vs.Definitiontext, &vs.Scopenotetext, &vs.Assigningauthorityid, &vs.Legacyflag, &vs.Statusdate)
		if err != nil {
			return nil, err
		}
		valueSets = append(valueSets, vs)
	}
	return &valueSets, nil
}

func GetValueSetByLikeOID(ctx context.Context, db xo.DB, oid string) (*[]xo.ValueSet, error) {
	wildcard := "%" + oid + "%"
	const sqlstr = "SELECT * FROM public.value_set WHERE oid LIKE $1"

	return queryValueSets(ctx, db, sqlstr, wildcard)
}

func GetValueSetByLikeID(ctx context.Context, db xo.DB, id string) (*[]xo.ValueSet, error) {
	wildcard := "%" + id + "%"
	const sqlstr = "SELECT * FROM public.value_set WHERE id LIKE $1"

	return queryValueSets(ctx, db, sqlstr, wildcard)
}

func GetValueSetByLikeString(ctx context.Context, db xo.DB, input string) (*[]xo.ValueSet, error) {
	wildcard := "%" + input + "%"
	const sqlstr = "SELECT * FROM public.value_set WHERE lower(name) LIKE lower($1) OR lower(code) LIKE lower($1);"

	return queryValueSets(ctx, db, sqlstr, wildcard)
}

func queryValueSets(ctx context.Context, db xo.DB, sqlstr, wildcard string) (*[]xo.ValueSet, error) {
	valueSets := []xo.ValueSet{}
	rows, err := db.QueryContext(ctx, sqlstr, wildcard)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		vs := xo.ValueSet{}
		err := rows.Scan(&vs.Oid, &vs.ID, &vs.Name, &vs.Code, &vs.Status, &vs.Definitiontext, &vs.Scopenotetext, &vs.Assigningauthorityid, &vs.Legacyflag, &vs.Statusdate)
		if err != nil {
			return nil, err
		}
		valueSets = append(valueSets, vs)
	}
	return &valueSets, nil
}
