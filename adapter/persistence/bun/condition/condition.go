package condition

import (
	"fmt"
	"strings"

	"github.com/ming-0x0/yuan/internal/common/repository"
	"github.com/uptrace/bun"
)

type Condition interface {
	repository.Condition
	Select(query *bun.SelectQuery) *bun.SelectQuery
	Delete(query *bun.DeleteQuery) *bun.DeleteQuery
	Update(query *bun.UpdateQuery) *bun.UpdateQuery
}

type relationCondition struct {
	relation string
}

func (c *relationCondition) Select(query *bun.SelectQuery) *bun.SelectQuery {
	return query.Relation(c.relation)
}

func (c *relationCondition) Delete(query *bun.DeleteQuery) *bun.DeleteQuery {
	return query
}

func (c *relationCondition) Update(query *bun.UpdateQuery) *bun.UpdateQuery {
	return query
}

func Relation(relation string) Condition {
	return &relationCondition{relation: relation}
}

type whereCondition struct {
	column   string
	operator string
	value    any
}

func (c *whereCondition) Select(query *bun.SelectQuery) *bun.SelectQuery {
	return query.Where("? "+c.operator+" ?", bun.Ident(c.column), c.value)
}

func (c *whereCondition) Delete(query *bun.DeleteQuery) *bun.DeleteQuery {
	return query.Where("? "+c.operator+" ?", bun.Ident(c.column), c.value)
}

func (c *whereCondition) Update(query *bun.UpdateQuery) *bun.UpdateQuery {
	return query.Where("? "+c.operator+" ?", bun.Ident(c.column), c.value)
}

func EQ(column string, value any) Condition {
	return &whereCondition{column: column, operator: "=", value: value}
}

func NEQ(column string, value any) Condition {
	return &whereCondition{column: column, operator: "!=", value: value}
}

func LIKE(column string, value any) Condition {
	return &whereCondition{column: column, operator: "ILIKE", value: "%" + strings.TrimSpace(fmt.Sprint(value)) + "%"}
}

type orCondition struct {
	conditions []Condition
}

func (c *orCondition) Select(query *bun.SelectQuery) *bun.SelectQuery {
	return query.WhereGroup(" AND ", func(q *bun.SelectQuery) *bun.SelectQuery {
		for _, condition := range c.conditions {
			q = q.WhereGroup(" OR ", condition.Select)
		}
		return q
	})
}

func (c *orCondition) Delete(query *bun.DeleteQuery) *bun.DeleteQuery {
	return query.WhereGroup(" AND ", func(q *bun.DeleteQuery) *bun.DeleteQuery {
		for _, condition := range c.conditions {
			q = q.WhereGroup(" OR ", condition.Delete)
		}
		return q
	})
}

func (c *orCondition) Update(query *bun.UpdateQuery) *bun.UpdateQuery {
	return query.WhereGroup(" AND ", func(q *bun.UpdateQuery) *bun.UpdateQuery {
		for _, condition := range c.conditions {
			q = q.WhereGroup(" OR ", condition.Update)
		}
		return q
	})
}

func OR(conditions ...Condition) Condition {
	return &orCondition{conditions: conditions}
}
