package xo

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

// HotTopic represents a row from 'public.hot_topic'.
type HotTopic struct {
	HotTopicID          uuid.UUID    `json:"hot_topic_id"`          // hot_topic_id
	HotTopicName        string       `json:"hot_topic_name"`        // hot_topic_name
	HotTopicDescription string       `json:"hot_topic_description"` // hot_topic_description
	Seq                 int          `json:"seq"`                   // seq
	EffectiveDate       time.Time    `json:"effective_date"`        // effective_date
	ExpiryDate          sql.NullTime `json:"expiry_date"`           // expiry_date
	DeploymentDate      time.Time    `json:"deployment_date"`       // deployment_date
	CreatedDate         time.Time    `json:"created_date"`          // created_date
	UpdatedDate         time.Time    `json:"updated_date"`          // updated_date
	// xo fields
	_exists, _deleted bool
}

// Exists returns true when the [HotTopic] exists in the database.
func (ht *HotTopic) Exists() bool {
	return ht._exists
}

// Deleted returns true when the [HotTopic] has been marked for deletion
// from the database.
func (ht *HotTopic) Deleted() bool {
	return ht._deleted
}

// Insert inserts the [HotTopic] to the database.
func (ht *HotTopic) Insert(ctx context.Context, db DB) error {
	switch {
	case ht._exists: // already exists
		return logerror(&ErrInsertFailed{ErrAlreadyExists})
	case ht._deleted: // deleted
		return logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
	// insert (manual)
	const sqlstr = `INSERT INTO public.hot_topic (` +
		`hot_topic_id, hot_topic_name, hot_topic_description, seq, effective_date, expiry_date, deployment_date, created_date, updated_date` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9` +
		`)`
	// run
	logf(sqlstr, ht.HotTopicID, ht.HotTopicName, ht.HotTopicDescription, ht.Seq, ht.EffectiveDate, ht.ExpiryDate, ht.DeploymentDate, ht.CreatedDate, ht.UpdatedDate)
	if _, err := db.ExecContext(ctx, sqlstr, ht.HotTopicID, ht.HotTopicName, ht.HotTopicDescription, ht.Seq, ht.EffectiveDate, ht.ExpiryDate, ht.DeploymentDate, ht.CreatedDate, ht.UpdatedDate); err != nil {
		return logerror(err)
	}
	// set exists
	ht._exists = true
	return nil
}

// Update updates a [HotTopic] in the database.
func (ht *HotTopic) Update(ctx context.Context, db DB) error {
	switch {
	case !ht._exists: // doesn't exist
		return logerror(&ErrUpdateFailed{ErrDoesNotExist})
	case ht._deleted: // deleted
		return logerror(&ErrUpdateFailed{ErrMarkedForDeletion})
	}
	// update with composite primary key
	const sqlstr = `UPDATE public.hot_topic SET ` +
		`hot_topic_name = $1, hot_topic_description = $2, seq = $3, effective_date = $4, expiry_date = $5, deployment_date = $6, created_date = $7, updated_date = $8 ` +
		`WHERE hot_topic_id = $9`
	// run
	logf(sqlstr, ht.HotTopicName, ht.HotTopicDescription, ht.Seq, ht.EffectiveDate, ht.ExpiryDate, ht.DeploymentDate, ht.CreatedDate, ht.UpdatedDate, ht.HotTopicID)
	if _, err := db.ExecContext(ctx, sqlstr, ht.HotTopicName, ht.HotTopicDescription, ht.Seq, ht.EffectiveDate, ht.ExpiryDate, ht.DeploymentDate, ht.CreatedDate, ht.UpdatedDate, ht.HotTopicID); err != nil {
		return logerror(err)
	}
	return nil
}

// Save saves the [HotTopic] to the database.
func (ht *HotTopic) Save(ctx context.Context, db DB) error {
	if ht.Exists() {
		return ht.Update(ctx, db)
	}
	return ht.Insert(ctx, db)
}

// Upsert performs an upsert for [HotTopic].
func (ht *HotTopic) Upsert(ctx context.Context, db DB) error {
	switch {
	case ht._deleted: // deleted
		return logerror(&ErrUpsertFailed{ErrMarkedForDeletion})
	}
	// upsert
	const sqlstr = `INSERT INTO public.hot_topic (` +
		`hot_topic_id, hot_topic_name, hot_topic_description, seq, effective_date, expiry_date, deployment_date, created_date, updated_date` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9` +
		`)` +
		` ON CONFLICT (hot_topic_id) DO ` +
		`UPDATE SET ` +
		`hot_topic_name = EXCLUDED.hot_topic_name, hot_topic_description = EXCLUDED.hot_topic_description, seq = EXCLUDED.seq, effective_date = EXCLUDED.effective_date, expiry_date = EXCLUDED.expiry_date, deployment_date = EXCLUDED.deployment_date, created_date = EXCLUDED.created_date, updated_date = EXCLUDED.updated_date `
	// run
	logf(sqlstr, ht.HotTopicID, ht.HotTopicName, ht.HotTopicDescription, ht.Seq, ht.EffectiveDate, ht.ExpiryDate, ht.DeploymentDate, ht.CreatedDate, ht.UpdatedDate)
	if _, err := db.ExecContext(ctx, sqlstr, ht.HotTopicID, ht.HotTopicName, ht.HotTopicDescription, ht.Seq, ht.EffectiveDate, ht.ExpiryDate, ht.DeploymentDate, ht.CreatedDate, ht.UpdatedDate); err != nil {
		return logerror(err)
	}
	// set exists
	ht._exists = true
	return nil
}

// Delete deletes the [HotTopic] from the database.
func (ht *HotTopic) Delete(ctx context.Context, db DB) error {
	switch {
	case !ht._exists: // doesn't exist
		return nil
	case ht._deleted: // deleted
		return nil
	}
	// delete with single primary key
	const sqlstr = `DELETE FROM public.hot_topic ` +
		`WHERE hot_topic_id = $1`
	// run
	logf(sqlstr, ht.HotTopicID)
	if _, err := db.ExecContext(ctx, sqlstr, ht.HotTopicID); err != nil {
		return logerror(err)
	}
	// set deleted
	ht._deleted = true
	return nil
}

// HotTopicByHotTopicID retrieves a row from 'public.hot_topic' as a [HotTopic].
//
// Generated from index 'hot_topic_pkey'.
func HotTopicByHotTopicID(ctx context.Context, db DB, hotTopicID uuid.UUID) (*HotTopic, error) {
	// query
	const sqlstr = `SELECT ` +
		`hot_topic_id, hot_topic_name, hot_topic_description, seq, effective_date, expiry_date, deployment_date, created_date, updated_date ` +
		`FROM public.hot_topic ` +
		`WHERE hot_topic_id = $1`
	// run
	logf(sqlstr, hotTopicID)
	ht := HotTopic{
		_exists: true,
	}
	if err := db.QueryRowContext(ctx, sqlstr, hotTopicID).Scan(&ht.HotTopicID, &ht.HotTopicName, &ht.HotTopicDescription, &ht.Seq, &ht.EffectiveDate, &ht.ExpiryDate, &ht.DeploymentDate, &ht.CreatedDate, &ht.UpdatedDate); err != nil {
		return nil, logerror(err)
	}
	return &ht, nil
}