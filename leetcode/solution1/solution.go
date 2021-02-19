package solution1

func twoSum(nums []int, target int) []int {
	index := make(map[int]int, len(nums))

	for i, x := range nums {
		if j, ok := index[target-x]; ok {
			return []int{j, i}
		}
		index[x] = i
	}

	return nil
}
