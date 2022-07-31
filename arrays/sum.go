package main

func Sum(numbers []int) (sum int) {
	for _, number := range numbers {
		sum += number
	}

	return
}

func SumAll(numbersToSum ...[]int) []int {
	lengthOfNumber := len(numbersToSum)
	sums := make([]int, lengthOfNumber)

	for i, number := range numbersToSum {
		sums[i] = Sum(number)
	}

	return sums
}
