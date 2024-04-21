package pqb_test

import (
	"fmt"

	"github.com/fritzkeyzer/pqb"
)

func ExampleSelect_Simple() {
	sql, args := pqb.Select().
		From("person").
		Where("name = $1", "John").
		Build()

	fmt.Println(sql)
	fmt.Println(args)
	// Output:
	// SELECT * FROM person WHERE name = $1;
	// [John]
}

func ExampleSelect_Conditions() {
	sql, args := pqb.Select().
		From("person").
		Conditions(
			pqb.Cond("name = $1", "John"),
			pqb.Cond("age > $1", 64),
			pqb.AnyOf(
				pqb.Cond("$1 = any(sports)", "tennis"),
				pqb.Cond("$1 = any(clubs)", "bowls"),
			),
		).
		Build()

	fmt.Println(sql)
	fmt.Println(args)
	// Output:
	// SELECT * FROM person WHERE (name = $1) AND (age > $2) AND (($3 = any(sports)) OR ($4 = any(clubs)));
	// [John 64 tennis bowls]
}

func ExampleSelectQuery_GroupBy() {
	sql, args := pqb.Select().
		From("person").
		GroupBy("age").
		OrderBy("name").
		Limit(10).
		Offset(5).
		Build()

	fmt.Println(sql)
	fmt.Println(args)
	// Output:
	// SELECT * FROM person GROUP BY age ORDER BY name LIMIT 10 OFFSET 5;
	// []
}

func ExampleSelectQuery_FromInner() {
	inner := pqb.Select().
		From("person").
		Where("name = $1", "John")

	sql, args := pqb.Select().
		Cols("count(*)").
		FromInner(inner).
		Build()

	fmt.Println(sql)
	fmt.Println(args)
	// Output:
	// SELECT count(*) FROM (SELECT * FROM person WHERE name = $1);
	// [John]
}

func ExampleSelectQuery_From() {
	sql, args := pqb.Select().
		Cols("count(*)").
		From("person inner join school on person.school_id = school.id").
		Build()

	fmt.Println(sql)
	fmt.Println(args)
	// Output:
	// SELECT count(*) FROM person inner join school on person.school_id = school.id;
	// []
}
