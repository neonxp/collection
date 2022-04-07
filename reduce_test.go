package collection

import (
	"reflect"
	"testing"
)

func TestReduce(t *testing.T) {
	type args struct {
		collection  []int
		cb          func(previous int, current int, idx int) int
		accumulator int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Sum",
			args: args{
				collection: []int{1, 2, 3, 4, 5, 6},
				cb: func(previous, current, idx int) int {
					return previous + current
				},
			},
			want: 21,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Reduce(tt.args.collection, tt.args.cb, tt.args.accumulator); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Reduce() = %v, want %v", got, tt.want)
			}
		})
	}
}
