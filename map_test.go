package collection

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMapSync(t *testing.T) {
	type args struct {
		collection []int
		cb         func(int, int) int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "multiple",
			args: args{
				collection: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				cb: func(item int, idx int) int {
					return item * idx
				},
			},
			want: []int{0, 2, 6, 12, 20, 30, 42, 56, 72, 90},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MapSync(tt.args.collection, tt.args.cb); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Map() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMap(t *testing.T) {
	type args struct {
		collection []int
		cb         func(int, int) string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "counter",
			args: args{
				collection: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				cb: func(item int, idx int) string {
					return fmt.Sprintf("%dth element is %d", idx, item)
				},
			},
			want: []string{
				"0th element is 1",
				"1th element is 2",
				"2th element is 3",
				"3th element is 4",
				"4th element is 5",
				"5th element is 6",
				"6th element is 7",
				"7th element is 8",
				"8th element is 9",
				"9th element is 10",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Map(tt.args.collection, tt.args.cb); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Map() = %v, want %v", got, tt.want)
			}
		})
	}
}
