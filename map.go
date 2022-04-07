package collection

import "sync"

func MapSync[T any, R any](collection []T, cb func(item T, idx int) R) []R {
	result := make([]R, len(collection))
	for i, v := range collection {
		result[i] = cb(v, i)
	}
	return result
}

func Map[T any, R any](collection []T, cb func(item T, idx int) R) []R {
	result := make([]R, len(collection))
	wg := sync.WaitGroup{}
	wg.Add(len(collection))
	for i, v := range collection {
		func(v T, i int) {
			go func() {
				defer wg.Done()
				result[i] = cb(v, i)
			}()
		}(v, i)
	}
	wg.Wait()
	return result
}
