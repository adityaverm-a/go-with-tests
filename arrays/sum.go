package main

func Sum(numbers []int) (sum int) {
	for _, number := range numbers {
		sum += number
	}

	return
}

func SumAll(numbersToSum ...[]int) (sums []int) {

	for _, number := range numbersToSum {
		sums = append(sums, Sum(number))
	}

	return
}

func SumAllTails(numbersToSum ...[]int) (sums []int) {

	for _, number := range numbersToSum {
		if len(number) == 0 {
			sums = append(sums, 0)
		} else {
			tail := number[1:]
			sums = append(sums, Sum(tail))
		}
	}

	return
}

// Quick Summary
// Things covered

// - Arrays
// - Slices
// 	- The various ways to make them
// 	- How they have a fixed capacity but you can create new slices from old ones using append
// 	- How to slice, slices!
// - len to get the length of an array or slice
// - Test coverage tool, `go test -cover`
// - reflect.DeepEqual and why it's useful but can reduce the type-safety of your code

// Used slices and arrays with integers but they work with any other type too, including arrays/slices themselves.
// So you can declare a variable of [][]string if you need to.
