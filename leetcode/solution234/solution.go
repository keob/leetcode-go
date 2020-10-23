package solution234

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	slow, fast := head, head
	var prev *ListNode = nil

	for fast != nil && fast.Next != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}

	prev.Next = nil
	var _head *ListNode = nil

	for slow != nil {
		tmp := slow.Next
		slow.Next = _head
		_head = slow
		slow = tmp
	}

	for head != nil && _head != nil {
		if head.Val != _head.Val {
			return false
		}
		head = head.Next
		_head = _head.Next
	}

	return true
}
