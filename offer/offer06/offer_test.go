package offer06

import (
	"reflect"
	"testing"
)

func ints2List(nums []int) *ListNode {
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

func TestReversePrint(t *testing.T) {
	tests := []struct {
		name string
		list []int
		want []int
	}{
		{"test 1", []int{1, 3, 2}, []int{2, 3, 1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := ints2List(tt.list)
			got := reversePrint(head)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reversePrint() = %v, want %v", got, tt.want)
			}
		})
	}
}
