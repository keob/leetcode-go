package utils

import (
	"reflect"
	"testing"
)

func TestList2Ints(t *testing.T) {
	list := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
			},
		},
	}

	want := []int{1, 2, 3}
	got := List2Ints(list)
	if !reflect.DeepEqual(want, got) {
		t.Errorf("expected '%v' but got '%v'", want, got)
	}
}

func TestInts2List(t *testing.T) {
	nums := []int{1, 2, 3}
	list := Ints2List(nums)
	got := []int{}

	for list != nil {
		got = append(got, list.Val)
		list = list.Next
	}

	if !reflect.DeepEqual(nums, got) {
		t.Errorf("expected '%v' but got '%v'", nums, got)
	}
}
