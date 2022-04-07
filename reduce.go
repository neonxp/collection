package collection

func Reduce[T any, R any](collection []T, cb func(previous R, current T, idx int) R, accumulator R) R {
	for i, v := range collection {
		accumulator = cb(accumulator, v, i)
	}
	return accumulator
}
