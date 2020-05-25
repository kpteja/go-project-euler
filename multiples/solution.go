package multiples

// CalculateSumOfMultiples calculates sum of multiples of passed numbers.
func CalculateSumOfMultiples(multiplesOf []int, from int, to int) int {
	sum := 0

	for from < to {
		for _, num := range multiplesOf {
			if from%num == 0 {
				sum = sum + from
				break
			}
		}

		from++
	}

	return sum
}
