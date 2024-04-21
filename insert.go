package pqb

type InsertQuery struct {
	table          string
	rows           Rows
	where          *Condition
	conflict       string
	conflictDoSql  string
	conflictDoArgs []any
	returning      string
	returningArgs  []any
}

// InsertInto creates a new InsertQuery. You will need to chain other methods
// onto this to build the query.
func InsertInto(table string) *InsertQuery {
	return &InsertQuery{
		table: table,
	}
}

// Vals sets the row to insert. Use this, ValsOf, Rows, or RowsFrom but not at the same time.
func (q *InsertQuery) Vals(values ...Value) *InsertQuery {
	q.rows = Rows{
		Rows: []Row{
			{
				Values: values,
			},
		},
	}
	return q
}

// Row sets the row to insert from a struct.
// (Use Vals, Row or Rows but not at the same time)
func (q *InsertQuery) Row(obj any, omitColumns ...string) *InsertQuery {
	q.rows = Rows{
		Rows: []Row{
			RowFrom(obj, omitColumns...),
		},
	}
	return q
}

// Rows creates a batch insert from a slice of structs.
// (Use Vals, Row or Rows but not at the same time)
func (q *InsertQuery) Rows(rows any) *InsertQuery {
	q.rows = RowsFrom(rows)
	return q
}

// Where sets the where clause. Use this or Filter but not both.
func (q *InsertQuery) Where(sql string, args ...any) *InsertQuery {
	cond := Cond(sql, args...)
	q.where = &cond
	return q
}

// Filter sets the where clause from a Condition. Use this or Where but not both.
func (q *InsertQuery) Filter(cond Condition) *InsertQuery {
	q.where = &cond
	return q
}

// OnConflictDoNothing allows you to specify a conflict target. Eg: "id"
func (q *InsertQuery) OnConflictDoNothing(conflict string) *InsertQuery {
	q.conflict = conflict
	q.conflictDoSql = "DO NOTHING"
	return q
}

// OnConflict allows you to specify a conflict target. Eg: "id"
func (q *InsertQuery) OnConflict(conflict string) *InsertQuery {
	q.conflict = conflict
	return q
}

// OnConflictDo allows you to specify the conflict action and arguments.
// Make sure the sql string begins with "DO".
// Be sure to specify the conflict target with OnConflict first.
func (q *InsertQuery) OnConflictDo(sql string, args ...any) *InsertQuery {
	q.conflictDoSql = sql
	q.conflictDoArgs = args

	return q
}

func (q *InsertQuery) Returning(sql string, args ...any) *InsertQuery {
	q.returning = sql
	q.returningArgs = args
	return q
}

// Build returns the SQL query and arguments
func (q *InsertQuery) Build() (string, []any) {
	var args []any

	colsSql, valsSql, valsArgs := q.rows.Insert()
	colsSql = "(" + IncrementArgs(colsSql, len(args)) + ")"
	valsSql = " VALUES " + IncrementArgs(valsSql, len(args))
	args = append(args, valsArgs...)

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

	conflict := ""
	if q.conflict != "" {
		conflict = " ON CONFLICT (" + q.conflict + ") " + IncrementArgs(q.conflictDoSql, len(args))
		args = append(args, q.conflictDoArgs...)
	}

	sql := "INSERT INTO " + q.table + colsSql + valsSql + where + conflict + returning + ";"

	return sql, args
}
