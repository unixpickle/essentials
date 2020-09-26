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
// Note that g may be called concurrently from multiple
// Goroutines at once.
//
// This can be useful if Goroutines each have their own
// local set of resources that they can reuse.
// For example, each Goroutine might have a connection
// pool or a local random number generator instance.
func StatefulConcurrentMap(maxGos, n int, g func() func(i int)) {
	ReduceConcurrentMap(maxGos, n, func() (func(int), func()) {
		return g(), nil
	})
}

// ReduceConcurrentMap is like StatefulConcurrentMap, but
// a final reduction function is called at the end of each
// Goroutine.
//
// The reduce function is called from one Goroutine at a
// time to allow aggregation operations to be unsafe.
// If the reduction function is nil, this is equivalent to
// StatefulConcurrencyMap.
//
// This can be used to have each Goroutine accumulate some
// partial information which is then aggregated.
func ReduceConcurrentMap(maxGos, n int, g func() (iter func(i int), reduce func())) {
	if maxGos == 0 {
		maxGos = runtime.GOMAXPROCS(0)
	}
	if maxGos > n {
		maxGos = n
	}

	var wg sync.WaitGroup
	var lock sync.Mutex
	for i := 0; i < maxGos; i++ {
		wg.Add(1)
		go func(start int) {
			defer wg.Done()
			f, reduce := g()
			for i := start; i < n; i += maxGos {
				f(i)
			}
			if reduce != nil {
				lock.Lock()
				defer lock.Unlock()
				reduce()
			}
		}(i)
	}
	wg.Wait()
}
