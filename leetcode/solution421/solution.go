package solution421

const highBits = 30

type trie struct {
	left  *trie
	right *trie
}

func (t *trie) add(num int) {
	cur := t
	for i := highBits; i >= 0; i-- {
		bit := num >> i & 1
		if bit == 0 {
			if cur.left == nil {
				cur.left = &trie{}
			}
			cur = cur.left
		} else {
			if cur.right == nil {
				cur.right = &trie{}
			}
			cur = cur.right
		}
	}
}

func (t *trie) check(num int) (x int) {
	cur := t
	for i := highBits; i >= 0; i-- {
		bit := num >> i & 1
		if bit == 0 {
			if cur.right != nil {
				cur = cur.right
				x = x*2 + 1
			} else {
				cur = cur.left
				x = x * 2
			}
		} else {
			if cur.left != nil {
				cur = cur.left
				x = x*2 + 1
			} else {
				cur = cur.right
				x = x * 2
			}
		}
	}
	return
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func findMaximumXOR(nums []int) (res int) {
	root := &trie{}

	for i := 1; i < len(nums); i++ {
		root.add(nums[i-1])
		res = max(res, root.check(nums[i]))
	}

	return
}
