package solution1052

func maxSatisfied(customers []int, grumpy []int, X int) int {
	sum := 0
	customersSize := len(customers)
	grumpySize := len(grumpy)

	for i := 0; i < grumpySize; i++ {
		if grumpy[i] == 0 {
			sum += customers[i]
			customers[i] = 0
		}
	}

	sumX := 0
	for i := 0; i < X; i++ {
		sumX += customers[i]
	}

	max := sumX
	for i := X; i < customersSize; i++ {
		sumX += customers[i] - customers[i-X]
		if sumX > max {
			max = sumX
		}
	}

	return max + sum
}
