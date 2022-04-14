package solution1672

func maximumWealth(accounts [][]int) (res int) {
	for _, account := range accounts {
		sum := 0
		for _, val := range account {
			sum += val
		}
		if sum >= res {
			res = sum
		}
	}
	return
}
