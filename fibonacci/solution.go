package fibonacci

// SumOfEvenValuedTerms calculates sum of even valued terms.
func SumOfEvenValuedTerms(maxSum uint64) uint64 {
	var prev uint64 = 1
	var cur uint64 = 2
	var sum uint64 = cur

	for sum < maxSum {
		next := prev + cur
		prev = cur
		cur = next

		if cur%2 == 0 {
			sum = sum + cur
		}
	}

	return sum
}
