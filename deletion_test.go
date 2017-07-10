package essentials

import "testing"

func TestUnorderedDelete(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5, 6}
	UnorderedDelete(&slice, 3)
	if !slicesEqual(slice, []int{1, 2, 3, 6, 5}) {
		t.Fatalf("unexpected slice: %v", slice)
	}
	UnorderedDelete(&slice, 4)
	if !slicesEqual(slice, []int{1, 2, 3, 6}) {
		t.Fatalf("unexpected slice: %v", slice)
	}
}

func TestOrderedDelete(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5, 6}
	OrderedDelete(&slice, 3)
	if !slicesEqual(slice, []int{1, 2, 3, 5, 6}) {
		t.Fatalf("unexpected slice: %v", slice)
	}
	OrderedDelete(&slice, 4)
	if !slicesEqual(slice, []int{1, 2, 3, 5}) {
		t.Fatalf("unexpected slice: %v", slice)
	}
}

func slicesEqual(s1, s2 []int) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i, x := range s1 {
		if s2[i] != x {
			return false
		}
	}
	return true
}
