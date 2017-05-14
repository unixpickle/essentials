package essentials

import "math"

// MaxInt computes the maximum of the arguments.
// If no arguments are provided, 0 is returned.
func MaxInt(ns ...int) int {
	if len(ns) == 0 {
		return 0
	}
	max := ns[0]
	for _, x := range ns[1:] {
		if x > max {
			max = x
		}
	}
	return max
}

// MinInt computes the minimum of the arguments.
// If no arguments are provided, 0 is returned.
func MinInt(ns ...int) int {
	if len(ns) == 0 {
		return 0
	}
	min := ns[0]
	for _, x := range ns[1:] {
		if x < min {
			min = x
		}
	}
	return min
}

// Round rounds to the nearest whole number.
//
// When x is the same distance from two different whole
// numbers, the one further from zero is selected.
func Round(x float64) float64 {
	if x < 0 {
		return math.Ceil(x - 0.5)
	} else {
		return math.Floor(x + 0.5)
	}
}
