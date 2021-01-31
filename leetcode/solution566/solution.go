package solution566

func matrixReshape(nums [][]int, r int, c int) [][]int {
	if len(nums) == 0 || len(nums[0]) == 0 {
		return nums
	}

	h := len(nums)
	l := len(nums[0])

	if h*l != r*c {
		return nums
	} else {
		var numsReshape = make([][]int, r)
		for i := 0; i < r; i++ {
			numsReshape[i] = make([]int, c)
		}
		var rcount = 0
		var ccount = 0

		for row := 0; row < h; row++ {
			for col := 0; col < l; col++ {
				numsReshape[rcount][ccount] = nums[row][col]
				ccount++
				if ccount == c {
					ccount = ccount - c
					rcount++
				}
			}
		}
		return numsReshape
	}
}
