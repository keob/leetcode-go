package solution1518

func numWaterBottles(numBottles int, numExchange int) int {
	if numBottles < numExchange {
		return numBottles
	}
	return (numBottles-numExchange)/(numExchange-1) + 1 + numBottles
}
