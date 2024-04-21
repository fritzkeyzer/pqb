package pqb_test

import (
	"fmt"

	"github.com/fritzkeyzer/pqb"
)

func ExampleVal() {
	_ = pqb.Val("name", "john")

	// is equivalent to:
	_ = pqb.Value{
		Col:  "name",
		Sql:  "$1",
		Args: []any{"john"},
	}

	// values can be used to build queries
	// if you need to use a custom sql fragment
	// you can use ValSql
	sql, args := pqb.InsertInto("person").Vals(
		pqb.Val("name", "John"),
		pqb.Val("city", "London"),
		pqb.Val("address", "123 Fake St"),
		pqb.ValSql("age", "age(now(), $1)", "1963-01-01"),
	).Build()

	fmt.Println("sql:", sql)
	fmt.Println("args:", args)
	// Output:
	// sql: INSERT INTO person(name, city, address, age) VALUES ($1, $2, $3, age(now(), $4));
	// args: [John London 123 Fake St 1963-01-01]
}

func ExampleValSql() {
	// values can be used to build queries
	// if you need to use a custom sql fragment
	// you can use ValSql
	sql, args := pqb.InsertInto("person").Vals(
		pqb.Val("name", "John"),
		pqb.Val("city", "London"),
		pqb.Val("address", "123 Fake St"),
		pqb.ValSql("age", "age(now(), $1)", "1963-01-01"),
	).Build()

	fmt.Println("sql:", sql)
	fmt.Println("args:", args)
	// Output:
	// sql: INSERT INTO person(name, city, address, age) VALUES ($1, $2, $3, age(now(), $4));
	// args: [John London 123 Fake St 1963-01-01]
}
