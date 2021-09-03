package solution61

import (
	"reflect"
	"testing"
)

func list2ints(head *ListNode) []int {
	res := []int{}

	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}

	return res
}

func ints2list(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	l := &ListNode{}
	list := l

	for _, v := range nums {
		list.Next = &ListNode{Val: v}
		list = list.Next
	}

	return l.Next
}

func TestRotateRight(t *testing.T) {
	tests := []struct {
		name string
		list []int
		k    int
		want []int
	}{
		{"test 1", []int{1, 2, 3, 4, 5}, 2, []int{4, 5, 1, 2, 3}},
		{"test 2", []int{0, 1, 2}, 4, []int{2, 0, 1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := ints2list(tt.list)
			res := rotateRight(head, tt.k)
			got := list2ints(res)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rotateRight() = %v, want %v", got, tt.want)
			}
		})
	}
}
