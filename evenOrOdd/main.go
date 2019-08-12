package main

import "fmt"

func main() {
	numbers := makeRange(0, 10)

	for i := range numbers {
		fmt.Println(evenOrOdd(numbers[i]));
	}
}

func makeRange(min, max int) []int {
	nums := make([]int, max-min+1)
	for i := range nums {
		nums[i] = min + i
	}

	return nums
}

func evenOrOdd(n int) (s string) {
	if n % 2 ==0 {
		s = fmt.Sprintf("%v is even.", n)
	} else {
		s = fmt.Sprintf("%v is odd.", n)
	}
	return
}
