package pqb

type UpsertQuery struct {
	table         string
	insertRow     Row
	conflict      string
	returning     string
	returningArgs []any
}

func Upsert(table string) *UpsertQuery {
	return &UpsertQuery{
		table: table,
	}
}

// Row sets the row to insert/update from a struct.
// (Use this or Vals but not at the same time)
func (q *UpsertQuery) Row(obj any, omit ...string) *UpsertQuery {
	q.insertRow = RowFrom(obj, omit...)
	return q
}

// Vals sets the values to insert/update from a struct.
// (Use this or Row but not at the same time)
func (q *UpsertQuery) Vals(vals ...Value) *UpsertQuery {
	q.insertRow = Row{
		Values: vals,
	}
	return q
}

// OnConflict allows you to specify a conflict target. Eg: "id"
func (q *UpsertQuery) OnConflict(conflict string) *UpsertQuery {
	q.conflict = conflict
	return q
}

func (q *UpsertQuery) Returning(sql string, args ...any) *UpsertQuery {
	q.returning = sql
	q.returningArgs = args
	return q
}

func (q *UpsertQuery) Build() (string, []any) {
	var args []any

	colsSql, valsSql, valsArgs := q.insertRow.Insert()
	colsSql = "(" + colsSql + ")"
	valsSql = " VALUES " + valsSql
	args = append(args, valsArgs...)

	updateSql, _ := q.insertRow.Update()

	onConflict := ""
	if q.conflict != "" {
		onConflict = " (" + q.conflict + ")"
	}

	returning := ""
	if q.returning != "" {
		returning = " RETURNING " + IncrementArgs(q.returning, len(args))
		args = append(args, q.returningArgs...)
	}

	sql := "INSERT INTO " + q.table + colsSql + valsSql +
		" ON CONFLICT" + onConflict + " DO UPDATE SET " + updateSql + returning + ";"

	return sql, args
}
