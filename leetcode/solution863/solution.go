package solution863

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(m map[*TreeNode]int, root *TreeNode, n int) {
	if root == nil {
		return
	}
	m[root] = n
	dfs(m, root.Left, n<<1+1)
	dfs(m, root.Right, n<<1+2)
}

func dist(n1, n2 int) (res int) {
	for n1 != n2 {
		if n1 > n2 {
			n1 = (n1 - 1) >> 1
		} else {
			n2 = (n2 - 1) >> 1
		}
		res++
	}
	return
}

func distanceK(root, target *TreeNode, k int) (res []int) {
	m := make(map[*TreeNode]int)
	dfs(m, root, 0)
	t := m[target]
	for n, v := range m {
		if dist(v, t) == k {
			res = append(res, n.Val)
		}
	}
	return
}
