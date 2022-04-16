package solution2235

func sum(num1 int, num2 int) int {
	for num2 != 0 {
		temp := num1
		num1 ^= num2
		num2 = (temp & num2) << 1
	}
	return num1
}
