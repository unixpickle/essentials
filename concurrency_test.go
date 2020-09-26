package essentials

import (
	"fmt"
	"strconv"
	"strings"
)

func ExampleConcurrentMap() {
	squares := make([]int, 20)
	ConcurrentMap(0, len(squares), func(i int) {
		squares[i] = i * i
	})
	fmt.Println(squares)

	// Output: [0 1 4 9 16 25 36 49 64 81 100 121 144 169 196 225 256 289 324 361]
}

func ExampleStatefulConcurrentMap() {
	isPrime := make([]bool, 70)
	StatefulConcurrentMap(0, len(isPrime), func() func(int) {
		// Each Goroutine has its own cache of primes.
		var primeCache []int

		return func(i int) {
			if i < 2 {
				// Don't test 0 and 1 for primality.
				return
			}
			// Speedup by checking if i is divisible by an existing
			// known prime.
			for _, p := range primeCache {
				if i%p == 0 {
					return
				}
			}
			// Brute force check.
			for j := 2; j < i; j++ {
				if i%j == 0 {
					primeCache = append(primeCache, j)
					return
				}
			}
			isPrime[i] = true
			primeCache = append(primeCache, i)
		}
	})

	// Print out the primes in a pretty way.
	var primes []string
	for i, x := range isPrime {
		if x {
			primes = append(primes, strconv.Itoa(i))
		}
	}
	fmt.Println(strings.Join(primes, ", "))

	// Output: 2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67
}

func ExampleReduceConcurrentMap() {
	// Compute the sum of integers from 0 to numSum - 1.
	numSum := 10000
	var totalSum int
	ReduceConcurrentMap(0, numSum, func() (func(int), func()) {
		var localSum int
		return func(i int) {
				// The iter function is not synchronized across Goroutines.
				localSum += i
			}, func() {
				// The reduce function is synchronized across Goroutines.
				totalSum += localSum
			}
	})
	fmt.Println(totalSum)

	// Output: 49995000
}
