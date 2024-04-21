package pqb_test

import (
	"fmt"

	"github.com/fritzkeyzer/pqb"
)

func ExampleCond() {
	filter := pqb.Cond("name = $1 and age > $2", "John", 18)

	sql, args := filter.Build()

	fmt.Println(sql)
	fmt.Println(args)
	// Output:
	// name = $1 and age > $2
	// [John 18]
}

func ExampleAnyOf() {
	filter := pqb.AnyOf(
		pqb.Cond("name = $1", "John"),
		pqb.Cond("name = $1", "Jane"),
	)

	sql, args := filter.Build()

	fmt.Println(sql)
	fmt.Println(args)
	// Output:
	// (name = $1) OR (name = $2)
	// [John Jane]
}

func ExampleAllOf() {
	filter := pqb.AllOf(
		pqb.Cond("name = $1", "John"),
		pqb.Cond("name = $1", "Jane"),
	)

	sql, args := filter.Build()

	fmt.Println(sql)
	fmt.Println(args)
	// Output:
	// (name = $1) AND (name = $2)
	// [John Jane]
}

func ExampleAllOf_nested() {
	filter := pqb.AllOf(
		pqb.Cond("a = $1 and b = $2", 2, 4),
		pqb.Cond("c = $1 and d = $2 and e = $3", 8, 16, 32),
		pqb.Cond("f = $1", 64),
		pqb.AllOf(
			pqb.Cond("g = $1", 128),
			pqb.AnyOf(
				pqb.Cond("h = $1", 256),
				pqb.Cond("i = $1 and j = $2", 512, 1024),
			),
		),
	)

	sql, args := filter.Build()

	fmt.Println(sql)
	fmt.Println(args)
	// Output:
	// (a = $1 and b = $2) AND (c = $3 and d = $4 and e = $5) AND (f = $6) AND ((g = $7) AND ((h = $8) OR (i = $9 and j = $10)))
	// [2 4 8 16 32 64 128 256 512 1024]
}

func ExampleNot() {
	filter := pqb.Not(
		pqb.Cond("name = $1", "John"),
	)

	sql, args := filter.Build()

	fmt.Println(sql)
	fmt.Println(args)
	// Output:
	// NOT (name = $1)
	// [John]
}

func ExampleNot_nested() {
	filter := pqb.AllOf(
		pqb.Cond("a = $1", 2),
		pqb.Not(
			pqb.AnyOf(
				pqb.Cond("b = $1", 4),
				pqb.Cond("c = $1 and d = $2", 8, 16),
			),
		),
	)

	sql, args := filter.Build()

	fmt.Println(sql)
	fmt.Println(args)
	// Output:
	// (a = $1) AND (NOT ((b = $2) OR (c = $3 and d = $4)))
	// [2 4 8 16]
}

func ExampleCondition_Build() {
	filter := pqb.AnyOf(
		pqb.Cond("name = $1", "John"),
		pqb.Cond("name = $1", "Jane"),
	)

	sql, args := filter.Build()

	fmt.Println(sql)
	fmt.Println(args)
	// Output:
	// (name = $1) OR (name = $2)
	// [John Jane]
}

func ExampleCondition_and() {
	filter := pqb.Cond("published = true")

	searchText := "Nature"
	if searchText != "" {
		filter.And(pqb.Cond("name ilike $1", searchText))
	}

	searchDate := "2020-01-01"
	if searchDate != "" {
		filter.And(pqb.Cond("published_at = date($1)", searchDate))
	}

	category := "art"
	if category != "" {
		filter.And(pqb.Cond("category = $1", category))
	}

	sql, args := filter.Build()

	fmt.Println(sql)
	fmt.Println(args)
	// Output:
	// (((published = true) AND (name ilike $1)) AND (published_at = date($2))) AND (category = $3)
	// [Nature 2020-01-01 art]
}
