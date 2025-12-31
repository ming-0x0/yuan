package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/ming-0x0/yuan/adapter/persistence/bun/condition"
	"github.com/ming-0x0/yuan/internal/common/apperror"
	"github.com/ming-0x0/yuan/internal/common/repository"
	ctxkey "github.com/ming-0x0/yuan/pkg/ctxutil/key"
	"github.com/ming-0x0/yuan/pkg/logger"
	"github.com/uptrace/bun"
)

type Model interface {
	repository.Model
}

type Converter[D, M any] interface {
	repository.Converter[D, M]
}

type RepositoryInterface[D any] interface {
	repository.Repository[D]
}

type Repository[C Converter[D, M], D, M any] struct {
	db        *bun.DB
	logger    *logger.Logger
	converter C
}

func New[C Converter[D, M], D, M any](
	db *bun.DB,
	logger *logger.Logger,
	converter C,
) *Repository[C, D, M] {
	return &Repository[C, D, M]{
		db:        db,
		logger:    logger,
		converter: converter,
	}
}

func (r *Repository[C, D, M]) DB(ctx context.Context) bun.IDB {
	v := ctx.Value(ctxkey.TransactionContextKey)
	if v != nil {
		if tx, ok := v.(bun.Tx); ok {
			return tx
		}
	}
	return r.db
}

func (r *Repository[C, D, M]) Create(
	ctx context.Context,
	domain *D,
	createdBy int64,
) error {
	model, err := r.converter.ToModel(domain)
	if err != nil {
		return err
	}

	if m, ok := any(model).(Model); ok {
		m.Create(createdBy)
	}

	if _, err := r.DB(ctx).NewInsert().Model(model).Exec(ctx); err != nil {
		return apperror.WithCause(apperror.Internal, err)
	}

	return nil
}

func (r *Repository[C, D, M]) TakeByConditions(
	ctx context.Context,
	conditions ...repository.Condition,
) (*D, error) {
	model := new(M)

	err := r.DB(ctx).
		NewSelect().
		Model(model).
		Apply(selectConditions(toBunConditions(conditions))...).
		Scan(ctx)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, apperror.WithCause(apperror.NotFound, err)
		}
		return nil, apperror.WithCause(apperror.Internal, err)
	}

	return r.converter.ToDomain(model)
}

func (r *Repository[C, D, M]) FindByConditions(
	ctx context.Context,
	conditions ...repository.Condition,
) ([]*D, error) {
	models := make([]*M, 0)

	err := r.DB(ctx).
		NewSelect().
		Model(&models).
		Apply(selectConditions(toBunConditions(conditions))...).
		Scan(ctx)

	if err != nil {
		return nil, apperror.WithCause(apperror.Internal, err)
	}

	return r.converter.ToDomainSlice(models)
}

func (r *Repository[C, D, M]) Update(
	ctx context.Context,
	domain *D,
	updatedBy int64,
) error {
	model, err := r.converter.ToModel(domain)
	if err != nil {
		return err
	}

	if m, ok := any(model).(Model); ok {
		m.Update(updatedBy)
	}

	if _, err := r.DB(ctx).NewUpdate().Model(model).OmitZero().WherePK().Exec(ctx); err != nil {
		return apperror.WithCause(apperror.Internal, err)
	}

	return nil
}

// Delete performs soft delete if supported, otherwise hard delete
func (r *Repository[C, D, M]) Delete(
	ctx context.Context,
	domain *D,
	deletedBy int64,
) error {
	model, err := r.converter.ToModel(domain)
	if err != nil {
		return err
	}

	// Soft delete if model supports it
	if _, ok := any(model).(Model); ok {
		return r.softDelete(ctx, model, deletedBy)
	}

	// Hard delete
	return r.hardDelete(ctx, model)
}

// softDelete performs soft delete using update
func (r *Repository[C, D, M]) softDelete(
	ctx context.Context,
	model *M,
	deletedBy int64,
) error {
	if m, ok := any(model).(Model); ok {
		m.Delete(deletedBy)
	}

	_, err := r.DB(ctx).
		NewUpdate().
		Model(model).
		Column("deleted_by", "deleted_at").
		WherePK().
		Exec(ctx)
	if err != nil {
		return apperror.WithCause(apperror.Internal, err)
	}

	return nil
}

// hardDelete performs permanent delete
func (r *Repository[C, D, M]) hardDelete(
	ctx context.Context,
	model *M,
) error {
	_, err := r.DB(ctx).
		NewDelete().
		Model(model).
		WherePK().
		Exec(ctx)
	if err != nil {
		return apperror.WithCause(apperror.Internal, err)
	}

	return nil
}

// toBunConditions converts generic conditions to bun-specific conditions
func toBunConditions(conditions []repository.Condition) []condition.Condition {
	bunConditions := make([]condition.Condition, 0, len(conditions))
	for _, cond := range conditions {
		if bc, ok := cond.(condition.Condition); ok {
			bunConditions = append(bunConditions, bc)
		}
	}
	return bunConditions
}

// selectConditions converts conditions to select query functions
func selectConditions(conditions []condition.Condition) []func(*bun.SelectQuery) *bun.SelectQuery {
	if len(conditions) == 0 {
		return nil
	}

	queries := make([]func(*bun.SelectQuery) *bun.SelectQuery, len(conditions))
	for i, condition := range conditions {
		queries[i] = condition.Select
	}
	return queries
}

// applyUpdateConditions converts conditions to update query functions
func updateConditions(conditions []condition.Condition) []func(*bun.UpdateQuery) *bun.UpdateQuery { //nolint:unused
	if len(conditions) == 0 {
		return nil
	}

	queries := make([]func(*bun.UpdateQuery) *bun.UpdateQuery, len(conditions))
	for i, condition := range conditions {
		queries[i] = condition.Update
	}
	return queries
}

// deleteConditions converts conditions to delete query functions
func deleteConditions(conditions []condition.Condition) []func(*bun.DeleteQuery) *bun.DeleteQuery { //nolint:unused
	if len(conditions) == 0 {
		return nil
	}

	queries := make([]func(*bun.DeleteQuery) *bun.DeleteQuery, len(conditions))
	for i, condition := range conditions {
		queries[i] = condition.Delete
	}
	return queries
}
