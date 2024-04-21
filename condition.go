package pqb

// Condition represents a SQL condition that would be used in a WHERE clause.
type Condition struct {
	sql  string
	args []any
}

// Cond creates a new Condition from the given SQL and args.
func Cond(sql string, args ...any) Condition {
	return Condition{
		sql:  sql,
		args: args,
	}
}

// AllOf creates a new Condition that is the logical AND of all the given Conditions.
func AllOf(conditions ...Condition) Condition {
	sql := ""
	var args []any

	for i := range conditions {
		if i > 0 {
			sql += " AND "
		}

		condSql := IncrementArgs(conditions[i].sql, len(args))

		sql += "(" + condSql + ")"
		args = append(args, conditions[i].args...)
	}

	return Condition{
		sql:  sql,
		args: args,
	}
}

// AnyOf creates a new Condition that is the logical OR of all the given Conditions.
func AnyOf(conditions ...Condition) Condition {
	sql := ""
	var args []any

	for i := range conditions {
		if i > 0 {
			sql += " OR "
		}

		condSql := IncrementArgs(conditions[i].sql, len(args))

		sql += "(" + condSql + ")"
		args = append(args, conditions[i].args...)
	}

	return Condition{
		sql:  sql,
		args: args,
	}
}

// Not creates a new Condition that is the logical NOT of the given Condition.
func Not(cond Condition) Condition {
	return Condition{
		sql:  "NOT (" + cond.sql + ")",
		args: cond.args,
	}
}

// And creates a new Condition that is the logical AND of the receiver and the provided Condition.
// The receiver is modified in place and is returned - allowing chaining.
// It is recommended to use AllOf instead of this method in most cases.
func (c *Condition) And(cond Condition) *Condition {
	condSql := IncrementArgs(cond.sql, len(c.args))

	if c.sql != "" {
		// while this is correct - it gets ugly wrapping brackets if there are many ANDs
		c.sql = "(" + c.sql + ") AND (" + condSql + ")"
	} else {
		// nothing to wrap, condition was empty, now it's not
		c.sql = condSql
	}

	c.args = append(c.args, cond.args...)
	return c
}

// Build returns the SQL string and arguments for the Condition.
func (c *Condition) Build() (string, []any) {
	return c.sql, c.args
}
