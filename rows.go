package pqb

import (
	"reflect"
	"slices"
	"strings"
)

type Rows struct {
	Rows []Row
}

func RowsFrom(slice any, omit ...string) Rows {
	v := reflect.ValueOf(slice)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Slice {
		return Rows{}
	}

	var ret []Row
	for i := 0; i < v.Len(); i++ {
		ret = append(ret, RowFrom(v.Index(i).Interface(), omit...))
	}
	return Rows{Rows: ret}
}

func (r Rows) InsertInto(table string) *InsertQuery {
	return &InsertQuery{
		table: table,
		rows:  r,
	}
}

func (r Rows) Insert() (string, string, []any) {
	// find all cols
	var cols []string
	for i := range r.Rows {
		for j := range r.Rows[i].Values {
			col := r.Rows[i].Values[j].Col

			if !slices.Contains(cols, col) {
				cols = append(cols, col)
			}
		}
	}

	// build up colsSql
	colsSql := strings.Join(cols, ", ")

	// build up valsSql and args for all rows
	valsSql := ""
	var args []any
	for i := range r.Rows {
		if i > 0 {
			valsSql += ", "
		}

		// build up row valsSql
		rowSql := ""
		for j := range cols {
			if j > 0 {
				rowSql += ", "
			}

			valSql := "NULL"

			// find matching col in row
			for k := range r.Rows[i].Values {
				if r.Rows[i].Values[k].Col == cols[j] {
					sql := r.Rows[i].Values[k].Sql
					valSql = IncrementArgs(sql, len(args))
					args = append(args, r.Rows[i].Values[k].Args...)
					break
				}
			}

			rowSql += valSql
		}

		valsSql += "(" + rowSql + ")"
	}

	return colsSql, valsSql, args
}
