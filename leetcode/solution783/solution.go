package solution783

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDiffInBST(root *TreeNode) int {
	var res []int
	stack := make([]*TreeNode, 0)
	min := 1<<63 - 1

	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, node.Val)
		if len(res) > 1 {
			r := node.Val - res[len(res)-2]
			if min > r {
				min = r
			}
		}
		root = node.Right
	}

	if len(res) < 2 {
		return res[0]
	}

	return min
}
