package solution1707

type Node struct {
	Left  *Node
	Right *Node
}

type Tree struct {
	root *Node
}

func (tree *Tree) insert(num int) {
	cur := tree.root
	for i := 31; i >= 0; i-- {
		bit := num & (1 << i)
		if bit == 0 {
			if cur.Left == nil {
				cur.Left = &Node{}
			}
			cur = cur.Left
		} else {
			if cur.Right == nil {
				cur.Right = &Node{}
			}
			cur = cur.Right
		}
	}
}

func maximizeXor(nums []int, queries [][]int) []int {
	tree := &Tree{}
	tree.root = &Node{}

	for _, num := range nums {
		tree.insert(num)
	}

	res := make([]int, len(queries))

	for i := 0; i < len(queries); i++ {
		res[i] = -1
		flag := false
		dfs(tree.root, queries[i][0], queries[i][1], 0, 31, i, &res, &flag)
	}

	return res
}

func dfs(cur *Node, base, limit, sum, deep, index int, res *[]int, flag *bool) {
	if sum > limit || cur == nil {
		return
	}

	if deep == -1 {
		(*res)[index] = sum ^ base
		*flag = true
		return
	}

	bit := base & (1 << deep)

	if bit == 0 {
		dfs(cur.Right, base, limit, sum+(1<<deep), deep-1, index, res, flag)
		if !(*flag) {
			dfs(cur.Left, base, limit, sum, deep-1, index, res, flag)
		}
	} else {
		dfs(cur.Left, base, limit, sum, deep-1, index, res, flag)
		if !(*flag) {
			dfs(cur.Right, base, limit, sum+(1<<deep), deep-1, index, res, flag)
		}
	}
}
