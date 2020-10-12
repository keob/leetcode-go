package solution530

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getMinimumDifference(root *TreeNode) int {
	if root == nil {
		return 0
	}

	min, preVal, cur := math.MaxInt64, -1, root
	updateMin := func(node *TreeNode, num int) {
		if preVal != -1 && node.Val-preVal < min {
			min = node.Val - preVal
		}
		preVal = node.Val
	}

	for cur != nil {
		if cur.Left == nil {
			updateMin(cur, preVal)
			cur = cur.Right
			continue
		}
		preNode := cur.Left
		for preNode.Right != nil && preNode.Right != cur {
			preNode = preNode.Right
		}
		if preNode.Right == nil {
			preNode.Right = cur
			cur = cur.Left
		} else {
			preNode.Right = nil
			updateMin(cur, preVal)
			cur = cur.Right
		}
	}

	return min
}
