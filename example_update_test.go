package pqb_test

import (
	"fmt"

	"github.com/fritzkeyzer/pqb"
)

func ExampleUpdate() {
	u := pqb.Update("person").
		Set("name", "John").
		Set("age", 18).
		Where("id = $1", 1)

	sql, args := u.Build()

	fmt.Println(sql)
	fmt.Println(args)
	// Output:
	// UPDATE person SET name = $1, age = $2 WHERE id = $3;
	// [John 18 1]
}

func ExampleUpdate_RowFrom() {
	type Person struct {
		Name string
		Age  int
	}

	john := Person{
		Name: "John",
		Age:  18,
	}

	u := pqb.Update("person").
		Row(john).
		Where("id = $1", 1)

	sql, args := u.Build()

	fmt.Println(sql)
	fmt.Println(args)
	// Output:
	// UPDATE person SET name = $1, age = $2 WHERE id = $3;
	// [John 18 1]
}

func ExampleUpdate_SetSql() {
	u := pqb.Update("person").
		Set("name", "John").
		SetSql("full_name", "$1 || $2", "John", "Doe").
		Where("id = $1", 1)

	sql, args := u.Build()

	fmt.Println(sql)
	fmt.Println(args)
	// Output:
	// UPDATE person SET name = $1, full_name = $2 || $3 WHERE id = $4;
	// [John John Doe 1]
}

func ExampleUpdate_Returning() {
	type Person struct {
		Name string
		Age  int
	}

	john := Person{
		Name: "John",
		Age:  18,
	}

	u := pqb.Update("person").
		Row(john).
		Where("id = $1", 1).
		Returning("now()")

	sql, args := u.Build()

	fmt.Println(sql)
	fmt.Println(args)
	// Output:
	// UPDATE person SET name = $1, age = $2 WHERE id = $3 RETURNING now();
	// [John 18 1]
}
