// Package utils of code utils
package utils

// ListNode is list node
type ListNode struct {
	Val  int
	Next *ListNode
}

// List2Ints list to array
func List2Ints(head *ListNode) []int {
	res := []int{}

	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}

	return res
}

// Ints2List array to list
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

// Ints2CList returns a list whose tail point to pos-indexed node
func Ints2CList(nums []int, pos int) *ListNode {
	head := Ints2List(nums)

	if pos == -1 {
		return head
	}

	c := head
	for pos > 0 {
		c = c.Next
		pos--
	}

	tail := c
	for tail.Next != nil {
		tail = tail.Next
	}

	tail.Next = c

	return head
}

// GetNodeWith returns the first node with val
func (list *ListNode) GetNodeWith(val int) *ListNode {
	res := list

	for res != nil {
		if res.Val == val {
			break
		}
		res = res.Next
	}

	return res
}
