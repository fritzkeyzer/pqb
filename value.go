package pqb

// Value represents a cell in a table.
type Value struct {
	Col  string // column name, eg: "name", "description", etc
	Sql  string // sql placeholder, eg: $1, $2, $3, etc
	Args []any  // value
}

// Val creates a Value from a key value pair. Eg: Val("score", 10)
// If you need to use a SQL expression (with or without arguments) - use ValSql instead.
func Val(col string, val any) Value {
	return Value{
		Col:  col,
		Sql:  "$1",
		Args: []any{val},
	}
}

// ValSql creates a Value from a column name, sql placeholder and arguments.
// In most cases, Val can be used instead.
// eg:
//
//	ValSql("score", "($1 + $2) / $3", 10, 20, 3)
func ValSql(col, sql string, args ...any) Value {
	return Value{
		Col:  col,
		Sql:  sql,
		Args: args,
	}
}
