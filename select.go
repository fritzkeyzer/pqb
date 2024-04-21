package pqb

import (
	"fmt"
	"strings"
)

type SelectQuery struct {
	with      string
	withArgs  []any
	cols      string
	colsArgs  []any
	from      string
	fromArgs  []any
	where     *Condition
	group     string
	groupArgs []any
	order     string
	orderArgs []any
	limit     *int
	offset    *int
}

// Select creates a new SelectQuery. It defaults to "SELECT *" from no table.
// You will need to chain other methods to build the query.
func Select() *SelectQuery {
	return &SelectQuery{
		cols: "*",
	}
}

// With sets the columns to select.
// (Use this or ColsOf but not at the same time)
func (q *SelectQuery) With(sql string, args ...any) *SelectQuery {
	q.with = sql
	q.withArgs = args
	return q
}

// Cols sets the columns to select.
// (Use this or ColsOf but not at the same time)
func (q *SelectQuery) Cols(sql string, args ...any) *SelectQuery {
	q.cols = sql
	q.colsArgs = args
	return q
}

// ColsOf sets the columns to select from the fields of a struct.
// (Use this or Cols but not at the same time)
func (q *SelectQuery) ColsOf(typ any, omit ...string) *SelectQuery {
	vals := RowFrom(typ, omit...).Values

	var cols []string
	for i := range vals {
		cols = append(cols, vals[i].Col)
	}

	q.cols = strings.Join(cols, ", ")
	q.colsArgs = nil

	return q
}

// From sets the from clause.
// (Use this or FromInner but not at the same time)
func (q *SelectQuery) From(fromSql string, args ...any) *SelectQuery {
	q.from = fromSql
	q.fromArgs = args
	return q
}

// FromInner sets the from clause to a subquery.
// (Use this or From but not at the same time)
func (q *SelectQuery) FromInner(subquery *SelectQuery) *SelectQuery {
	sql, args := subquery.Build()
	sql = strings.TrimSuffix(sql, ";")

	q.from = "(" + sql + ")"
	q.fromArgs = args

	return q
}

// Where sets the where clause.
// (Use this or Conditions but not at the same time)
func (q *SelectQuery) Where(sql string, args ...any) *SelectQuery {
	cond := Cond(sql, args...)
	q.where = &cond
	return q
}

// Conditions sets the where clause to the logical AND of the given conditions.
// (Use this or Where but not at the same time)
func (q *SelectQuery) Conditions(conditions ...Condition) *SelectQuery {
	cond := AllOf(conditions...)
	q.where = &cond
	return q
}

func (q *SelectQuery) GroupBy(sql string, args ...any) *SelectQuery {
	q.group = sql
	q.groupArgs = args
	return q
}

func (q *SelectQuery) OrderBy(sql string, args ...any) *SelectQuery {
	q.order = sql
	q.orderArgs = args
	return q
}

func (q *SelectQuery) Limit(limit int) *SelectQuery {
	q.limit = &limit
	return q
}

func (q *SelectQuery) Offset(offset int) *SelectQuery {
	q.offset = &offset
	return q
}

// Build builds the query and returns the SQL and args.
func (q *SelectQuery) Build() (string, []any) {
	args := q.colsArgs

	with := ""
	if q.with != "" {
		with = IncrementArgs(q.with, len(args)) + " "
		args = append(args, q.withArgs...)
	}

	from := ""
	if q.from != "" {
		from = " FROM " + IncrementArgs(q.from, len(args))
		args = append(args, q.fromArgs...)
	}

	where := ""
	if q.where != nil {
		var whereArgs []any
		where, whereArgs = q.where.Build()
		where = " WHERE " + IncrementArgs(where, len(args))

		args = append(args, whereArgs...)
	}

	group := ""
	if q.group != "" {
		group = " GROUP BY " + IncrementArgs(q.group, len(args))
		args = append(args, q.groupArgs...)
	}

	order := ""
	if q.order != "" {
		order = " ORDER BY " + IncrementArgs(q.order, len(args))
		args = append(args, q.orderArgs...)
	}

	limit := ""
	offset := ""
	if q.limit != nil {
		limit = fmt.Sprintf(" LIMIT %d", *q.limit)

		if q.offset != nil && *q.offset > 0 {
			offset = fmt.Sprintf(" OFFSET %d", *q.offset)
		}
	}

	sql := with + "SELECT " + q.cols + from + where + group + order + limit + offset + ";"

	return sql, args
}
