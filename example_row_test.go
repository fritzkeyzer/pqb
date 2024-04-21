package pqb_test

import (
	"fmt"

	"github.com/fritzkeyzer/pqb"
)

func ExampleRowFrom() {
	type User struct {
		Id   int
		Name string
	}

	user := User{
		Id:   1,
		Name: "John",
	}

	row := pqb.RowFrom(user)

	for _, value := range row.Values {
		fmt.Printf("%#v\n", value)
	}

	// Output:
	// pqb.Value{Col:"id", Sql:"$1", Args:[]interface {}{1}}
	// pqb.Value{Col:"name", Sql:"$1", Args:[]interface {}{"John"}}
}

func ExampleRow_Insert() {
	type User struct {
		Id   int
		Name string
	}

	user := User{
		Id:   1,
		Name: "John",
	}

	cols, vals, args := pqb.RowFrom(user).Insert()

	fmt.Println("cols:", cols)
	fmt.Println("vals:", vals)
	fmt.Println("args:", args)

	// Output:
	// cols: id, name
	// vals: ($1, $2)
	// args: [1 John]
}

func ExampleRow_Update() {
	type User struct {
		Id   int
		Name string
	}

	user := User{
		Id:   1,
		Name: "John",
	}

	sql, args := pqb.RowFrom(user).Update()

	fmt.Println("sql:", sql)
	fmt.Println("args:", args)
	// Output:
	// sql: id = $1, name = $2
	// args: [1 John]
}

func ExampleRow_InsertInto() {
	type User struct {
		Id   int
		Name string
	}

	user := User{
		Id:   1,
		Name: "John",
	}

	sql, args := pqb.RowFrom(user).
		InsertInto("users").
		Build()

	fmt.Println("sql:", sql)
	fmt.Println("args:", args)
	// Output:
	// sql: INSERT INTO users(id, name) VALUES ($1, $2);
	// args: [1 John]
}

func ExampleRow_UpdateTable() {
	type User struct {
		Id   int
		Name string
	}

	user := User{
		Id:   1,
		Name: "John",
	}

	sql, args := pqb.RowFrom(user, "id").
		UpdateTable("users").
		Where("id = $1", 1).
		Build()

	fmt.Println("sql:", sql)
	fmt.Println("args:", args)
	// Output:
	// sql: UPDATE users SET name = $1 WHERE id = $2;
	// args: [John 1]
}
