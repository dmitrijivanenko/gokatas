package pfactor

func Factor(number int) []int {
	var result = make([]int, 0)

	if number == 1 {
		return []int{1}
	}

	for prime := 2; number > 1; prime++ {
		for number%prime == 0 {
			number = number / prime
			result = append(result, prime)
		}
	}

	return result
}
