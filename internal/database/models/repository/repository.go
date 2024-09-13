package repository

import (
	"context"
	"database/sql"

	"github.com/CDCgov/phinvads-fhir/internal/database/models"
	"github.com/CDCgov/phinvads-fhir/internal/database/models/xo"
)

// Repository contains all the CRUD methods for all resource types represented in the models package
type Repository struct {
	database xo.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{database: db}
}

// =============================== //
// ====== CodeSystem methods ===== //
// =============================== //
func (r *Repository) GetAllCodeSystems(ctx context.Context) (*[]xo.CodeSystem, error) {
	return models.GetAllCodeSystems(ctx, r.database)
}

func (r *Repository) GetCodeSystemByID(ctx context.Context, id string) (*xo.CodeSystem, error) {
	return xo.CodeSystemByID(ctx, r.database, id)
}

func (r *Repository) GetCodeSystemByOID(ctx context.Context, oid string) (*xo.CodeSystem, error) {
	return xo.CodeSystemByOid(ctx, r.database, oid)
}

// =============================== //
// == CodeSystemConcept methods == //
// =============================== //
func (r *Repository) GetAllCodeSystemConcepts(ctx context.Context) (*[]xo.CodeSystemConcept, error) {
	return models.GetAllCodeSystemConcepts(ctx, r.database)
}

func (r *Repository) GetCodeSystemConceptByID(ctx context.Context, id string) (*xo.CodeSystemConcept, error) {
	return xo.CodeSystemConceptByID(ctx, r.database, id)
}

func (r *Repository) GetCodeSystemConceptsByOID(ctx context.Context, oid string) ([]*xo.CodeSystemConcept, error) {
	result, err := xo.CodeSystemConceptByCodesystemoid(ctx, r.database, oid)
	if len(result) == 0 {
		err = sql.ErrNoRows
	}
	return result, err
}

func (r *Repository) GetCodeSystemByValueSetConceptCsOid(ctx context.Context, vsc *xo.ValueSetConcept) (*xo.CodeSystem, error) {
	return vsc.CodeSystem(ctx, r.database)
}

// =============================== //
// ====== ValueSet methods ======= //
// =============================== //
func (r *Repository) GetAllValueSets(ctx context.Context) (*[]xo.ValueSet, error) {
	return models.GetAllValueSets(ctx, r.database)
}

func (r *Repository) GetValueSetByID(ctx context.Context, id string) (*xo.ValueSet, error) {
	return xo.ValueSetByID(ctx, r.database, id)
}

func (r *Repository) GetValueSetByOID(ctx context.Context, oid string) (*xo.ValueSet, error) {
	return xo.ValueSetByOid(ctx, r.database, oid)
}

func (r *Repository) GetValueSetByVersionOID(ctx context.Context, vsv *xo.ValueSetVersion) (*xo.ValueSet, error) {
	return vsv.ValueSet(ctx, r.database)
}

// =============================== //
// ========= View methods ======== //
// =============================== //
func (r *Repository) GetAllViews(ctx context.Context) (*[]xo.View, error) {
	return models.GetAllViews(ctx, r.database)
}

func (r *Repository) GetViewByID(ctx context.Context, id string) (*xo.View, error) {
	return xo.ViewByID(ctx, r.database, id)
}

func (r *Repository) GetViewByViewVersionId(ctx context.Context, vv *xo.ViewVersion) (*xo.View, error) {
	return vv.View(ctx, r.database)
}

// =============================== //
// ===== ViewValueSet methods ==== //
// =============================== //
func (r *Repository) GetViewValueSetVersionByVvIdVsvId(ctx context.Context, viewVersionId, valueSetVersionId string) (*xo.ViewValueSetVersion, error) {
	return xo.ViewValueSetVersionByViewversionidValuesetversionid(ctx, r.database, viewVersionId, valueSetVersionId)
}

// =============================== //
// ===== ViewVersion methods ===== //
// =============================== //
func (r *Repository) GetViewVersionByID(ctx context.Context, id string) (*xo.ViewVersion, error) {
	return xo.ViewVersionByID(ctx, r.database, id)
}

func (r *Repository) GetViewVersionByViewId(ctx context.Context, viewId string) ([]*xo.ViewVersion, error) {
	result, err := xo.ViewVersionByViewid(ctx, r.database, viewId)
	if len(result) == 0 {
		err = sql.ErrNoRows
	}
	return result, err
}

func (r *Repository) GetViewVersionByVvsvVvId(ctx context.Context, vvsv *xo.ViewValueSetVersion) (*xo.ViewVersion, error) {
	return vvsv.ViewVersion(ctx, r.database)
}

// =============================== //
// === ValueSetConcept methods === //
// =============================== //
func (r *Repository) GetValueSetConceptsByCodeSystemOID(ctx context.Context, csOid string) ([]*xo.ValueSetConcept, error) {
	result, err := xo.ValueSetConceptByCodesystemoid(ctx, r.database, csOid)
	if len(result) == 0 {
		err = sql.ErrNoRows
	}
	return result, err
}

func (r *Repository) GetValueSetConceptByID(ctx context.Context, id string) (*xo.ValueSetConcept, error) {
	return xo.ValueSetConceptByID(ctx, r.database, id)
}

func (r *Repository) GetValueSetConceptByValueSetVersionID(ctx context.Context, vsvId string) ([]*xo.ValueSetConcept, error) {
	result, err := xo.ValueSetConceptByValuesetversionid(ctx, r.database, vsvId)
	if len(result) == 0 {
		err = sql.ErrNoRows
	}
	return result, err
}

// =============================== //
// ==== ValueSetGroup methods ==== //
// =============================== //
func (r *Repository) GetValueSetGroupByID(ctx context.Context, id string) (*xo.ValueSetGroup, error) {
	return xo.ValueSetGroupByID(ctx, r.database, id)
}

// =============================== //
// === ValueSetVersion methods === //
// =============================== //

func (r *Repository) GetValueSetVersionByID(ctx context.Context, id string) (*xo.ValueSetVersion, error) {
	return xo.ValueSetVersionByID(ctx, r.database, id)
}

func (r *Repository) GetValueSetVersionByValueSetOID(ctx context.Context, oid string) ([]*xo.ValueSetVersion, error) {
	result, err := xo.ValueSetVersionByValuesetoid(ctx, r.database, oid)
	if len(result) == 0 {
		err = sql.ErrNoRows
	}
	return result, err
}

func (r *Repository) GetValueSetVersionByVscVsvId(ctx context.Context, vsc *xo.ValueSetConcept) (*xo.ValueSetVersion, error) {
	return vsc.ValueSetVersion(ctx, r.database)
}

func (r *Repository) GetValueSetVersionByVvsvVsvId(ctx context.Context, vvsv *xo.ViewValueSetVersion) (*xo.ValueSetVersion, error) {
	return vvsv.ValueSetVersion(ctx, r.database)
}
