package offer06

type ListNode struct {
	Val  int
	Next *ListNode
}

func reversePrint(head *ListNode) []int {
	var (
		data []int
		res  []int
	)

	for head != nil {
		data = append(data, head.Val)
		head = head.Next
	}

	for len(data) != 0 {
		res = append(res, data[len(data)-1])
		data = data[:len(data)-1]
	}

	return res
}
