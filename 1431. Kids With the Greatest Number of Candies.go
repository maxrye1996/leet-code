package main

import "fmt"

func kidsWithCandies(candies []int, extraCandies int) []bool {
	biggest := candies[0]
	for _, num := range candies[1:] {
		if num > biggest {
			biggest = num
		}
	}

	result := []bool{}
	for _, num := range candies {
		if num+extraCandies >= biggest {
			result = append(result, true)
		} else {
			result = append(result, false)
		}

	}
	return result
}

func main() {
	fmt.Println(kidsWithCandies([]int{2, 3, 5, 1, 3}, 3))

	fmt.Println(kidsWithCandies([]int{4, 2, 1, 1, 2}, 1))

	fmt.Println(kidsWithCandies([]int{12, 1, 12}, 10))
}
