package solution61

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if k == 0 || head == nil || head.Next == nil {
		return head
	}

	n, cur := 1, head

	for cur.Next != nil {
		cur = cur.Next
		n++
	}

	cur.Next = head
	k %= n

	for i := 1; i <= n-k; i++ {
		cur = cur.Next
	}

	head, cur.Next = cur.Next, nil

	return head
}
