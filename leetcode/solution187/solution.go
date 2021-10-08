package solution187

const LEN int = 10

var m = map[byte]int{'A': 0, 'C': 1, 'G': 2, 'T': 3}

func findRepeatedDnaSequences(s string) (res []string) {
	n := len(s)

	if n <= LEN {
		return
	}

	x := 0

	for _, ch := range s[:LEN-1] {
		x = x<<2 | m[byte(ch)]
	}

	cnt := map[int]int{}

	for i := 0; i <= n-LEN; i++ {
		x = (x<<2 | m[s[i+LEN-1]]) & (1<<(LEN*2) - 1)
		cnt[x]++
		if cnt[x] == 2 {
			res = append(res, s[i:i+LEN])
		}
	}
	return res
}
