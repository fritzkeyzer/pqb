package pqb_test

import (
	"fmt"

	"github.com/fritzkeyzer/pqb"
)

func ExampleRows_Insert() {
	type User struct {
		Id   int
		Name string
	}

	users := []User{
		{
			Id:   1,
			Name: "John",
		},
		{
			Id:   2,
			Name: "Jane",
		},
		{
			Id:   3,
			Name: "Bob",
		},
	}

	cols, vals, args := pqb.RowsFrom(users).Insert()

	fmt.Println("cols:", cols)
	fmt.Println("vals:", vals)
	fmt.Println("args:", args)
	// Output:
	// cols: id, name
	// vals: ($1, $2), ($3, $4), ($5, $6)
	// args: [1 John 2 Jane 3 Bob]
}

func ExampleRows_InsertInto() {
	type User struct {
		Id   int
		Name string
	}

	users := []User{
		{
			Id:   1,
			Name: "John",
		},
		{
			Id:   2,
			Name: "Wayne",
		},
		{
			Id:   3,
			Name: "Bob",
		},
	}

	sql, args := pqb.RowsFrom(users).
		InsertInto("users").
		Build()

	fmt.Println("sql:", sql)
	fmt.Println("args:", args)
	// Output:
	// sql: INSERT INTO users(id, name) VALUES ($1, $2), ($3, $4), ($5, $6);
	// args: [1 John 2 Wayne 3 Bob]
}
