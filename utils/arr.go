package utils

// Flat function creates a new array
// with matrix elements concatenated into it recursively up to
func Flat(matrix [][]int) []int {
	res := []int{}

	for _, m := range matrix {
		res = append(res, m...)
	}

	return res
}
