package pqb

import (
	"reflect"
	"testing"
)

func TestSelectQuery_Build(t *testing.T) {
	type fields struct {
		cols      string
		colsArgs  []any
		from      string
		fromArgs  []any
		where     *Condition
		group     string
		groupArgs []any
		order     string
		orderArgs []any
		limit     *int
		offset    *int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
		want1  []any
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &SelectQuery{
				cols:      tt.fields.cols,
				colsArgs:  tt.fields.colsArgs,
				from:      tt.fields.from,
				fromArgs:  tt.fields.fromArgs,
				where:     tt.fields.where,
				group:     tt.fields.group,
				groupArgs: tt.fields.groupArgs,
				order:     tt.fields.order,
				orderArgs: tt.fields.orderArgs,
				limit:     tt.fields.limit,
				offset:    tt.fields.offset,
			}
			got, got1 := q.Build()
			if got != tt.want {
				t.Errorf("Build() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Build() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
