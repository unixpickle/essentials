package essentials

import (
	"runtime"
	"sync"
)

// ConcurrentMap calls f for every integer in [0, n).
//
// If maxGos is 0, then GOMAXPROCS goroutines are started.
// Otherwise, up to maxGos goroutines are used.
func ConcurrentMap(maxGos, n int, f func(i int)) {
	StatefulConcurrentMap(maxGos, n, func() func(int) {
		return f
	})
}

// StatefulConcurrentMap is like ConcurrentMap, but it
// calls g once per Goroutine, and then calls the result
// of g with every index on that Goroutine.
//
// This can be useful if Goroutines each have their own
// local set of resources that they can reuse.
func StatefulConcurrentMap(maxGos, n int, g func() func(i int)) {
	if maxGos == 0 {
		maxGos = runtime.GOMAXPROCS(0)
	}
	if maxGos > n {
		maxGos = n
	}

	var wg sync.WaitGroup
	for i := 0; i < maxGos; i++ {
		wg.Add(1)
		go func(start int) {
			defer wg.Done()
			f := g()
			for i := start; i < n; i += maxGos {
				f(i)
			}
		}(i)
	}
	wg.Wait()
}
