package collection

import (
	"reflect"
	"testing"
)

func TestPushPop(t *testing.T) {
	collection := []int{}
	collection = Push(collection, 1)
	collection = Push(collection, 2)
	collection = Push(collection, 3)
	want := []int{1, 2, 3}
	if !reflect.DeepEqual(collection, want) {
		t.Errorf("Want %+v, but got %+v", want, collection)
	}
	collection, e := Pop(collection)
	if e != 3 {
		t.Errorf("Want 3, but got %d", e)
	}
	collection, e = Pop(collection)
	if e != 2 {
		t.Errorf("Want 2, but got %d", e)
	}
	collection, e = Pop(collection)
	if e != 1 {
		t.Errorf("Want 1, but got %d", e)
	}
	collection, e = Pop(collection)
	if e != 0 {
		t.Errorf("Want 0, but got %d", e)
	}
	if len(collection) != 0 {
		t.Errorf("Collection must be empty, but got %+v (len = %d)", collection, len(collection))
	}
}
