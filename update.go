package pqb

type UpdateQuery struct {
	table         string
	values        []Value
	from          string
	fromArgs      []any
	where         *Condition
	returning     string
	returningArgs []any
}

// Update creates a new UpdateQuery. You will need to chain other methods onto
// this to build the query.
func Update(table string) *UpdateQuery {
	return &UpdateQuery{
		table: table,
	}
}

// Set appends a column-value pair. This can be chained multiple times and used
// in conjunction with SetSql.
func (q *UpdateQuery) Set(col string, val any) *UpdateQuery {
	q.values = append(q.values, Value{
		Col:  col,
		Sql:  "$1",
		Args: []any{val},
	})
	return q
}

// SetSql appends a column-value pair using a SQL expression. This can be chained multiple times and used
// in conjunction with Set.
func (q *UpdateQuery) SetSql(col string, sql string, args ...any) *UpdateQuery {
	q.values = append(q.values, Value{
		Col:  col,
		Sql:  sql,
		Args: args,
	})
	return q
}

// Row sets the values to set from a struct. Use this or Set / SetSql but not at the same time.
func (q *UpdateQuery) Row(obj any, omit ...string) *UpdateQuery {
	q.values = RowFrom(obj, omit...).Values
	return q
}

// Where sets the WHERE clause.
// (Use this or Conditions but not at the same time)
func (q *UpdateQuery) Where(sql string, args ...any) *UpdateQuery {
	cond := Cond(sql, args...)
	q.where = &cond
	return q
}

// Conditions sets the WHERE clause as the logical AND of the provided conditions.
// (Use this or Where but not at the same time)
func (q *UpdateQuery) Conditions(conditions ...Condition) *UpdateQuery {
	cond := AllOf(conditions...)
	q.where = &cond
	return q
}

func (q *UpdateQuery) Returning(sql string, args ...any) *UpdateQuery {
	q.returning = sql
	q.returningArgs = args
	return q
}

// Build returns the SQL query and arguments
func (q *UpdateQuery) Build() (string, []any) {
	setSql, args := Row{Values: q.values}.Update()

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

	returning := ""
	if q.returning != "" {
		returning = " RETURNING " + IncrementArgs(q.returning, len(args))
		args = append(args, q.returningArgs...)
	}

	sql := "UPDATE " + q.table + " SET " + setSql + from + where + returning + ";"

	return sql, args
}
