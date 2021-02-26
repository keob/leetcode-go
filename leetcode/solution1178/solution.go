package solution1178

import "sort"

type trieNode struct {
	son [26]*trieNode
	cnt int
}

func findNumOfValidWords(words []string, puzzles []string) []int {
	root := &trieNode{}
	for _, word := range words {
		w := []byte(word)
		sort.Slice(w, func(i, j int) bool {
			return w[i] < w[j]
		})

		i := 0
		for _, ch := range w[1:] {
			if w[i] != ch {
				i++
				w[i] = ch
			}
		}
		w = w[:i+1]

		cur := root
		for _, ch := range w {
			ch -= 'a'
			if cur.son[ch] == nil {
				cur.son[ch] = &trieNode{}
			}
			cur =cur.son[ch]
		}
		cur.cnt++
	}

	ans := make([]int, len(puzzles))
	for i, puzzle := range puzzles {
		pz := []byte(puzzle)
		first := pz[0]
		sort.Slice(pz, func(i, j int) bool {
			return pz[i] < pz[j]
		})

		var find func(int, *trieNode) int
		find = func(pos int, cur *trieNode) int {
			if cur == nil {
				return 0
			}

			if pos == len(pz) {
				return cur.cnt
			}

			res := find(pos +1, cur.son[ pz[pos]-'a' ])

			if pz[pos]!= first {
				res += find(pos+1, cur)
			}

			return res
		}

		ans[i] = find(0, root)
	}

	return ans
}
