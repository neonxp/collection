package collection

func Push[T any](collection []T, element T) []T {
	return append(collection, element)
}

func Pop[T any](collection []T) ([]T, T) {
	if len(collection) == 0 {
		return collection, *new(T)
	}
	return collection[:len(collection)-1], collection[len(collection)-1]
}
