package pqb

import (
	"fmt"
	"testing"
)

func ExampleIncrementArgs() {
	sql := "a = $99 and b = $2 and c = $52 and d = $4"

	sql = IncrementArgs(sql, 10)

	fmt.Println(sql)
	// Output:
	// a = $109 and b = $12 and c = $62 and d = $14
}

func TestIncrementArgs(t *testing.T) {
	type args struct {
		sql  string
		bump int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "none",
			args: args{
				sql:  "a = $1",
				bump: 0,
			},
			want: "a = $1",
		},
		{
			name: "simple",
			args: args{
				sql:  "a = $1",
				bump: 1,
			},
			want: "a = $2",
		},
		{
			name: "> 10",
			args: args{
				sql:  "a = $1 and b = $2 and c = $3 and d = $4 and e = $5 and f = $6 and g = $7 and h = $8 and i = $9 and j = $10 and k = $11",
				bump: 10,
			},
			want: "a = $11 and b = $12 and c = $13 and d = $14 and e = $15 and f = $16 and g = $17 and h = $18 and i = $19 and j = $20 and k = $21",
		},
		{
			name: "> 10 out of order",
			args: args{
				sql:  "a = $99 and b = $2 and c = $3 and d = $4 and e = $52 and f = $6 and g = $7 and h = $8 and i = $9 and j = $80 and k = $11",
				bump: 10,
			},
			want: "a = $109 and b = $12 and c = $13 and d = $14 and e = $62 and f = $16 and g = $17 and h = $18 and i = $19 and j = $90 and k = $21",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IncrementArgs(tt.args.sql, tt.args.bump); got != tt.want {
				t.Errorf("got:\n%#v\nwant:\n%#v\n", got, tt.want)
			}
		})
	}
}
