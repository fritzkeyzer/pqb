package pqb

import (
	"reflect"
	"strings"

	"github.com/stoewer/go-strcase"
)

type Row struct {
	Values []Value
}

type Valuer interface {
	SqlValue() (string, []any)
}

func RowFrom(obj any, omit ...string) Row {
	v := reflect.ValueOf(obj)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		return Row{}
	}

	t := v.Type()

	var values []Value
	for i := 0; i < t.NumField(); i++ {
		// skip unexported fields
		if !t.Field(i).IsExported() {
			continue
		}

		// get column name
		// check struct tags for "db" tag
		// if not found, use field name
		col := ""
		if tag, ok := t.Field(i).Tag.Lookup("db"); ok {
			col = tag
		} else {
			col = strcase.SnakeCase(t.Field(i).Name)
		}

		// skip column if it's in the omit list
		omitted := false
		for _, o := range omit {
			if col == o {
				omitted = true
				break
			}
		}
		if omitted {
			continue
		}

		val := v.Field(i).Interface()

		// skip nil values
		if val == nil {
			continue
		}

		if valuer, ok := val.(Valuer); ok {
			sql, args := valuer.SqlValue()
			values = append(values, Value{
				Col:  col,
				Sql:  sql,
				Args: args,
			})
		} else {
			values = append(values, Value{
				Col:  col,
				Sql:  "$1",
				Args: []any{val},
			})
		}

	}

	return Row{
		Values: values,
	}
}

func (r Row) Insert() (string, string, []any) {
	rows := Rows{Rows: []Row{r}}

	return rows.Insert()
}

func (r Row) Update() (string, []any) {
	var sets []string
	var args []any
	for i := range r.Values {
		set := r.Values[i].Col + " = " + r.Values[i].Sql
		sets = append(sets, IncrementArgs(set, len(args)))
		args = append(args, r.Values[i].Args...)
	}

	return strings.Join(sets, ", "), args
}

func (r Row) UpdateTable(table string) *UpdateQuery {
	return &UpdateQuery{
		table:  table,
		values: r.Values,
	}
}

func (r Row) InsertInto(table string) *InsertQuery {
	return &InsertQuery{
		table: table,
		rows:  Rows{Rows: []Row{r}},
	}
}
