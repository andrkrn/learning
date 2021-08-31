package arrays

func Sum(numbers []int) (sum int) {
	for _, number := range numbers {
		sum += number
	}

	return
}

func SumAll(numbers ...[]int) []int {
	var sum []int

	for _, n := range numbers {
		sum = append(sum, Sum(n))
	}

	return sum
}

func SumAllTails(numbers ...[]int) []int {
	var sum []int

	for _, n := range numbers {
		if len(n) == 0 {
			sum = append(sum, 0)
		} else {
			sum = append(sum, Sum(n[1:]))
		}
	}

	return sum
}
