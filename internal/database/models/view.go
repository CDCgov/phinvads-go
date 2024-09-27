package models

import (
	"context"

	"github.com/CDCgov/phinvads-go/internal/database/models/xo"
)

func GetAllViews(ctx context.Context, db xo.DB) (*[]xo.View, error) {
	const sqlstr = `SELECT * FROM public.view`
	views := []xo.View{}
	rows, err := db.QueryContext(ctx, sqlstr)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		v := xo.View{}
		err := rows.Scan(&v.Name, &v.ID, &v.Descriptiontext, &v.Status, &v.Viewnotes, &v.Statusdate)
		if err != nil {
			return nil, err
		}
		views = append(views, v)
	}
	return &views, nil
}
