package pqb_test

import (
	"reflect"
	"testing"

	"github.com/fritzkeyzer/pqb"
)

func ptr[T any](in T) *T {
	return &in
}

func TestRowFrom(t *testing.T) {
	type args struct {
		in   any
		omit []string
	}
	ptrId := ptr("an id")

	tests := []struct {
		name string
		args args
		want []pqb.Value
	}{
		{
			name: "string",
			args: args{
				in: "string",
			},
		},
		{
			name: "string ptr",
			args: args{
				in: ptr("string"),
			},
		},
		{
			name: "basic string and int",
			args: args{
				in: struct {
					Name string
					Age  int
				}{
					Name: "John",
					Age:  42,
				},
			},
			want: []pqb.Value{
				{
					Col:  "name",
					Sql:  "$1",
					Args: []any{"John"},
				},
				{
					Col:  "age",
					Sql:  "$1",
					Args: []any{42},
				},
			},
		},
		{
			name: "struct ptr",
			args: args{
				in: &struct {
					Name string
					Age  int
				}{
					Name: "John",
					Age:  42,
				},
			},
			want: []pqb.Value{
				pqb.Val("name", "John"),
				pqb.Val("age", 42),
			},
		},
		{
			name: "tag",
			args: args{
				in: struct {
					Name string `db:"custom_name"`
				}{
					Name: "John",
				},
			},
			want: []pqb.Value{
				pqb.Val("custom_name", "John"),
			},
		},
		{
			name: "snake case",
			args: args{
				in: struct {
					CamelCase string
					URLs      string `db:"urls"`
				}{
					CamelCase: "hello",
					URLs:      "world",
				},
			},
			want: []pqb.Value{
				pqb.Val("camel_case", "hello"),
				pqb.Val("urls", "world"),
			},
		},
		{
			name: "omit",
			args: args{
				in: struct {
					Id   string
					Name string
				}{
					Id:   "123",
					Name: "Freddy",
				},
				omit: []string{"id"},
			},
			want: []pqb.Value{
				pqb.Val("name", "Freddy"),
			},
		},
		{
			name: "pointers",
			args: args{
				in: struct {
					Id   *string
					Name *string
				}{
					Id:   ptrId,
					Name: nil,
				},
			},
			want: []pqb.Value{
				pqb.Val("id", ptrId),
				pqb.Val("name", (*string)(nil)),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := pqb.RowFrom(tt.args.in, tt.args.omit...).Values
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got: \n%#v\nwant:\n%#v", got, tt.want)
			}
		})
	}
}
