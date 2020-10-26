package solution1365

func smallerNumbersThanCurrent(nums []int) []int {
	cnt := [101]int{}
	ans := make([]int, len(nums))

	for _, v := range nums {
		cnt[v]++
	}

	for i := 0; i < 100; i++ {
		cnt[i+1] += cnt[i]
	}

	for k, v := range nums {
		if v > 0 {
			ans[k] = cnt[v-1]
		}
	}

	return ans
}
