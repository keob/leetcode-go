package solution725

type ListNode struct {
	Val  int
	Next *ListNode
}

func splitListToParts(head *ListNode, k int) []*ListNode {
	n := 0

	for node := head; node != nil; node = node.Next {
		n++
	}

	quotient, remainder := n/k, n%k
	parts := make([]*ListNode, k)

	for i, cur := 0, head; i < k && cur != nil; i++ {
		parts[i] = cur
		partSize := quotient
		if i < remainder {
			partSize++
		}
		for j := 1; j < partSize; j++ {
			cur = cur.Next
		}
		cur, cur.Next = cur.Next, nil
	}

	return parts
}
