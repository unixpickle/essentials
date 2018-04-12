package essentials

import (
	"reflect"
	"sort"
)

// Reverse reverses a slice in-place.
func Reverse(slice interface{}) {
	size := reflect.ValueOf(slice).Len()
	sw := swapper(slice)
	for i := 0; i < size/2; i++ {
		sw(i, size-i-1)
	}
}

// VoodooSort sorts the list using the comparator, while
// simultaneously re-ordering a set of other lists to
// match the re-ordering of the sorted list.
// In a sense, what is done to the sorted list is also
// done to the other lists, making the sorted list like a
// voodoo doll.
func VoodooSort(slice interface{}, less func(i, j int) bool, other ...interface{}) {
	vs := &voodooSorter{
		length:   reflect.ValueOf(slice).Len(),
		swappers: []func(i, j int){swapper(slice)},
		less:     less,
	}
	for _, o := range other {
		vs.swappers = append(vs.swappers, swapper(o))
	}
	sort.Sort(vs)
}

type voodooSorter struct {
	length   int
	swappers []func(i, j int)
	less     func(i, j int) bool
}

func (v *voodooSorter) Len() int {
	return v.length
}

func (v *voodooSorter) Swap(i, j int) {
	for _, s := range v.swappers {
		s(i, j)
	}
}

func (v *voodooSorter) Less(i, j int) bool {
	return v.less(i, j)
}
