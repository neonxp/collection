package collection

import "fmt"

func ExampleEach() {
	collection := []int{1, 2, 3, 4, 5, 6}
	Each(collection, func(v int, idx int) {
		fmt.Printf("Element %d: %d\n", idx, v)
	})
}

func ExampleEachSync() {
	collection := []int{1, 2, 3, 4, 5, 6}
	EachSync(collection, func(v int, idx int) {
		fmt.Printf("Element %d: %d\n", idx, v)
	})
}
