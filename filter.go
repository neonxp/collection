package collection

import "sync"

func FilterSync[T any](collection []T, filter func(item T, idx int) bool) []T {
	var result []T
	for i, v := range collection {
		if filter(v, i) {
			result = append(result, v)
		}
	}
	return result
}

func Filter[T any](collection []T, filter func(item T, idx int) bool) []T {
	var result []T
	mu := sync.Mutex{}
	wg := sync.WaitGroup{}
	wg.Add(len(collection))
	for i, v := range collection {
		func(v T, i int) {
			go func() {
				defer wg.Done()
				if filter(v, i) {
					mu.Lock()
					result = append(result, v)
					mu.Unlock()
				}
			}()
		}(v, i)
	}
	wg.Wait()
	return result
}
