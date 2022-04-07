package collection

import (
	"fmt"
	"strings"
)

func ExampleMap() {
	collection := []int{1, 2, 3, 4, 5}
	cb := func(v int, idx int) string {
		return fmt.Sprintf("[%d]", v)
	}
	result := Map(collection, cb)
	fmt.Println(strings.Join(result, "_"))
	// Output: [1]_[2]_[3]_[4]_[5]
}

func ExampleMapSync() {
	collection := []int{1, 2, 3, 4, 5}
	cb := func(v int, idx int) string {
		return fmt.Sprintf("[%d]", v)
	}
	result := MapSync(collection, cb)
	fmt.Println(strings.Join(result, "_"))
	// Output: [1]_[2]_[3]_[4]_[5]
}
