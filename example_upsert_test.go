package pqb_test

import (
	"fmt"

	"github.com/fritzkeyzer/pqb"
)

func ExampleUpsert() {
	type Person struct {
		Id   int
		Name string
		Age  int
	}

	john := Person{
		Id:   12,
		Name: "John",
		Age:  42,
	}

	sql, args := pqb.Upsert("person").
		Row(john, "id").
		Build()

	fmt.Println("sql:", sql)
	fmt.Println("args:", args)
	// Output:
	// sql: INSERT INTO person(name, age) VALUES ($1, $2) ON CONFLICT DO UPDATE SET name = $1, age = $2;
	// args: [John 42]
}

func ExampleUpsert_Conflict() {
	type Data struct {
		Id        int
		ProfileId int
		Source    string
		SourceId  string
		Name      string
		Content   string
	}

	raw := Data{
		Id:        12,
		ProfileId: 34,
		Source:    "twitter",
		SourceId:  "123",
		Name:      "John",
		Content:   "Hello!",
	}

	sql, args := pqb.Upsert("raw").
		Row(raw, "id", "profile_id").
		OnConflict("source, source_id").
		Returning("id, profile_id").
		Build()

	fmt.Println("sql:", sql)
	fmt.Println("args:", args)
	// Output:
	// sql: INSERT INTO raw(source, source_id, name, content) VALUES ($1, $2, $3, $4) ON CONFLICT (source, source_id) DO UPDATE SET source = $1, source_id = $2, name = $3, content = $4 RETURNING id, profile_id;
	// args: [twitter 123 John Hello!]
}
