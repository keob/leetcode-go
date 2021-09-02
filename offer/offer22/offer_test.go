package offer22

import (
	"reflect"
	"testing"
)

func List2Ints(head *ListNode) []int {
	res := []int{}

	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}

	return res
}

func Ints2List(nums []int) *ListNode {
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

func TestGetKthFromEnd(t *testing.T) {
	tests := []struct {
		name string
		head []int
		k    int
		want []int
	}{
		{"test 1", []int{1, 2, 3, 4, 5}, 2, []int{4, 5}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := Ints2List(tt.head)
			res := getKthFromEnd(list, tt.k)
			got := List2Ints(res)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getKthFromEnd() = %v, want %v", got, tt.want)
			}
		})
	}
}
