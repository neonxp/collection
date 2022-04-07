package collection

import (
	"reflect"
	"testing"
)

func TestFilterSync(t *testing.T) {
	type args struct {
		collection []int
		filter     func(item int, idx int) bool
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "odds",
			args: args{
				collection: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				filter: func(item int, idx int) bool {
					return item%2 == 0
				},
			},
			want: []int{2, 4, 6, 8, 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FilterSync(tt.args.collection, tt.args.filter); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Filter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilter(t *testing.T) {
	type args struct {
		collection []int
		filter     func(item int, idx int) bool
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "odds count",
			args: args{
				collection: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				filter: func(item int, idx int) bool {
					return item%2 == 0
				},
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Filter(tt.args.collection, tt.args.filter); len(got) != tt.want {
				t.Errorf("FilterParallel() returned %v elements, want %v", len(got), tt.want)
			}
		})
	}
}
