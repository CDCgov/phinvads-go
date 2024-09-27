package models

import (
	"context"

	"github.com/CDCgov/phinvads-go/internal/database/models/xo"
)

// All retrieves all rows from 'public.hot_topic'
func GetAllHotTopics(ctx context.Context, db xo.DB) (*[]xo.HotTopic, error) {
	const sqlstr = `SELECT * FROM public.hot_topic`
	hotTopics := []xo.HotTopic{}
	rows, err := db.QueryContext(ctx, sqlstr)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		ht := xo.HotTopic{}
		err := rows.Scan(&ht.HotTopicID, &ht.HotTopicName, &ht.HotTopicDescription, &ht.Seq, &ht.EffectiveDate, &ht.ExpiryDate, &ht.DeploymentDate, &ht.CreatedDate, &ht.UpdatedDate)
		if err != nil {
			return nil, err
		}
		hotTopics = append(hotTopics, ht)
	}
	return &hotTopics, nil
}
