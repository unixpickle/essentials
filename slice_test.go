package essentials

import (
	"reflect"
	"testing"
)

func TestReverse(t *testing.T) {
	pairs := [][2][]int{
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
		{[]int{1, 2}, []int{2, 1}},
		{[]int{1, 2, 3}, []int{3, 2, 1}},
		{[]int{1, 2, 3, 4}, []int{4, 3, 2, 1}},
		{[]int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}},
	}
	for _, pair := range pairs {
		old := append([]int{}, pair[0]...)
		Reverse(pair[0])
		if !reflect.DeepEqual(pair[0], pair[1]) {
			t.Errorf("reverse of %v should be %v but got %v", old, pair[1], pair[0])
		}
	}
}

func TestContains(t *testing.T) {
	list1 := []int{1, 2, 3}
	for _, val := range list1 {
		if !Contains(list1, val) {
			t.Errorf("did not contain %d", val)
		}
	}
	if Contains(list1, 4) {
		t.Error("should not contain 4")
	}

	list2 := []interface{}{"hey", nil, "test"}
	for _, val := range list2 {
		if !Contains(list2, val) {
			t.Errorf("did not contain %v", val)
		}
	}
	if Contains(list2, "yo") {
		t.Error("should not contain \"yo\"")
	}
}

func TestVoodooSort(t *testing.T) {
	list := []int{5, 3, 4, 7, 9, 1}
	other1 := []string{
		"fourth",
		"second",
		"third",
		"fifth",
		"sixth",
		"first",
	}
	other2 := make([]sillyStruct, len(list))
	for i, j := range list {
		other2[i] = sillyStruct{field1: uint64(i), field3: uint64(j)}
	}
	VoodooSort(list, func(i, j int) bool {
		return list[i] < list[j]
	}, other1, other2)

	expectedList := []int{1, 3, 4, 5, 7, 9}
	expectedOther1 := []string{"first", "second", "third", "fourth", "fifth", "sixth"}
	expectedOther2 := []sillyStruct{
		{5, 0, 1},
		{1, 0, 3},
		{2, 0, 4},
		{0, 0, 5},
		{3, 0, 7},
		{4, 0, 9},
	}
	if !reflect.DeepEqual(list, expectedList) {
		t.Errorf("list should be %v but got %v", expectedList, list)
	}
	if !reflect.DeepEqual(other1, expectedOther1) {
		t.Errorf("other 1 should be %v but got %v", expectedOther1, other1)
	}
	if !reflect.DeepEqual(other2, expectedOther2) {
		t.Errorf("other 2 should be %v but got %v", expectedOther2, other2)
	}
}

type sillyStruct struct {
	field1 uint64
	field2 uint64
	field3 uint64
}
