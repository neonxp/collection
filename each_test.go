package collection

import (
	"sync/atomic"
	"testing"
)

func TestEachSync(t *testing.T) {
	collection := []int{1, 2, 3, 4, 5, 6}
	want := 21
	sum := 0
	EachSync(collection, func(v int, _ int) {
		sum += v
	})
	if sum != want {
		t.Errorf("Expected %d, got %d", want, sum)
	}
}

func TestEach(t *testing.T) {
	collection := []int{1, 2, 3, 4, 5, 6}
	want := int64(21)
	sum := int64(0)
	Each(collection, func(v int, _ int) {
		atomic.AddInt64(&sum, int64(v))
	})
	if sum != want {
		t.Errorf("Expected %d, got %d", want, sum)
	}
}
