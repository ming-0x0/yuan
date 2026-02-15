package bun

import (
	"context"
	"database/sql"

	"github.com/ming-0x0/yuan/internal/modules/common/domain/id"
	"github.com/ming-0x0/yuan/internal/modules/iam/domain"
	"github.com/ming-0x0/yuan/internal/modules/iam/domain/permissiongroup"
	orm "github.com/ming-0x0/yuan/pkg/orm/bun/db"
)

type pgRepository struct {
	db *orm.DB
}

func NewPermissionGroupRepository(db *orm.DB) domain.PermissionGroupRepository {
	return &pgRepository{db: db}
}

func (r *pgRepository) FindByID(ctx context.Context, id id.ID) (*permissiongroup.PermissionGroup, error) {
	m := new(pgModel)
	err := r.db.WithContext(ctx).NewSelect().Model(m).Where("id = ?", id).Scan(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return pgToDomain(m), nil
}

func (r *pgRepository) ListAll(ctx context.Context) ([]*permissiongroup.PermissionGroup, error) {
	var ms []pgModel
	err := r.db.WithContext(ctx).NewSelect().Model(&ms).Scan(ctx)
	if err != nil {
		return nil, err
	}

	var pgs []*permissiongroup.PermissionGroup
	for _, m := range ms {
		pgs = append(pgs, pgToDomain(&m))
	}
	return pgs, nil
}
