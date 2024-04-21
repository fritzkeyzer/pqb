package pqb_test

import (
	"fmt"

	"github.com/fritzkeyzer/pqb"
)

func ExampleInsertInto() {
	ins := pqb.InsertInto("person").Vals(
		pqb.Val("name", "John"),
		pqb.Val("age", 18),
		pqb.Val("sex", "male"),
		pqb.Val("city", "Amsterdam"),
	)

	sql, args := ins.Build()

	fmt.Println(sql)
	fmt.Println(args)
	// Output:
	// INSERT INTO person(name, age, sex, city) VALUES ($1, $2, $3, $4);
	// [John 18 male Amsterdam]
}

func ExampleInsertInto_RowFrom() {
	type Person struct {
		Name string
		Age  int
	}

	john := Person{
		Name: "John",
		Age:  18,
	}

	ins := pqb.InsertInto("person").Row(john)

	sql, args := ins.Build()

	fmt.Println(sql)
	fmt.Println(args)
	// Output:
	// INSERT INTO person(name, age) VALUES ($1, $2);
	// [John 18]
}

func ExampleInsertInto_RowsFrom() {
	type Person struct {
		Name string
		Age  int
	}

	people := []Person{
		{Name: "John", Age: 18},
		{Name: "Jane", Age: 19},
		{Name: "Jack", Age: 20},
		{Name: "Jill", Age: 21},
	}

	ins := pqb.InsertInto("person").Rows(people)

	sql, args := ins.Build()

	fmt.Println(sql)
	fmt.Println(args)
	// Output:
	// INSERT INTO person(name, age) VALUES ($1, $2), ($3, $4), ($5, $6), ($7, $8);
	// [John 18 Jane 19 Jack 20 Jill 21]
}

func ExampleInsertInto_OnConflictDo() {
	sql, args := pqb.InsertInto("person").
		Vals(
			pqb.Val("name", "John"),
			pqb.Val("age", 18),
			pqb.Val("sex", "male"),
			pqb.Val("city", "Amsterdam"),
		).
		OnConflict("name, city").
		OnConflictDo(
			"do update set updated_at = now(), something = $1",
			"foo",
		).
		Returning("id").
		Build()

	fmt.Println(sql)
	fmt.Println(args)
	// Output:
	// INSERT INTO person(name, age, sex, city) VALUES ($1, $2, $3, $4) ON CONFLICT (name, city) do update set updated_at = now(), something = $5 RETURNING id;
	// [John 18 male Amsterdam foo]
}
