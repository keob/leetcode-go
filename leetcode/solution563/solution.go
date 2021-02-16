package solution563

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findTilt(root *TreeNode) (ans int) {
	var find func(*TreeNode) (int, int)
	find = func(node *TreeNode) (int, int) {
		if node == nil {
			return 0, 0
		}

		sumLeft, tiltLeft := find(node.Left)
		sumRight, tiltRight := find(node.Right)

		sum := sumLeft + sumRight + node.Val
		tilt := sumLeft - sumRight

		if tilt < 0 {
			tilt = -tilt
		}

		tilt += tiltLeft + tiltRight

		return sum, tilt
	}
	_, ans = find(root)

	return
}
