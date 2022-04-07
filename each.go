package collection

import "sync"

func EachSync[T any](collection []T, cb func(item T, idx int)) {
	for i, v := range collection {
		cb(v, i)
	}
}

func Each[T any](collection []T, cb func(item T, idx int)) {
	wg := sync.WaitGroup{}
	wg.Add(len(collection))
	for i, v := range collection {
		func(i int, v T) {
			go func() {
				defer wg.Done()
				cb(v, i)
			}()
		}(i, v)
	}
	wg.Wait()
}
