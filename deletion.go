package essentials

import "reflect"

// UnorderedDelete removes the indexed element from the
// slice, potentially changing the order of the slice in
// the process.
//
// The slicePtr argument must be a pointer to a slice.
//
// This performs the deletion in O(1) time, at the expense
// of re-ordering the set.
// For an order-preserving deletion, see OrderedDelete.
func UnorderedDelete(slicePtr interface{}, idx int) {
	slicePtrVal := reflect.ValueOf(slicePtr)
	if slicePtrVal.Type().Kind() != reflect.Ptr ||
		slicePtrVal.Type().Elem().Kind() != reflect.Slice {
		panic("first argument must be slice pointer")
	}
	slice := slicePtrVal.Elem()
	if idx < 0 || idx >= slice.Len() {
		panic("index out of range")
	}
	slice.Index(idx).Set(slice.Index(slice.Len() - 1))
	shrinkSlice(slice)
}

// OrderedDelete removes the indexed element from the
// slice and moves the following elements to fill its
// place.
//
// The slicePtr argument must be a pointer to a slice.
//
// This performs the deletion in O(N) time, with the
// benefit that it preserves the order of the slice.
// For a deletion that ignores order, see UnorderedDelete.
func OrderedDelete(slicePtr interface{}, idx int) {
	slicePtrVal := reflect.ValueOf(slicePtr)
	if slicePtrVal.Type().Kind() != reflect.Ptr ||
		slicePtrVal.Type().Elem().Kind() != reflect.Slice {
		panic("first argument must be slice pointer")
	}
	slice := slicePtrVal.Elem()
	if idx < 0 || idx >= slice.Len() {
		panic("index out of range")
	}
	reflect.Copy(slice.Slice(idx, slice.Len()-1), slice.Slice(idx+1, slice.Len()))
	shrinkSlice(slice)
}

func shrinkSlice(slice reflect.Value) {
	// Zero last element to prevent memory leak.
	lastElem := slice.Index(slice.Len() - 1)
	lastElem.Set(reflect.Zero(lastElem.Type()))

	slice.Set(slice.Slice(0, slice.Len()-1))
}
