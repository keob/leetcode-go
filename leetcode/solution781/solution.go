package solution781

func numRabbits(answers []int) (ans int) {
	count := map[int]int{}

	for _, v := range answers {
		count[v]++
	}

	for k, v := range count {
		ans += (k + v) / (k + 1) * (k + 1)
	}

	return
}
