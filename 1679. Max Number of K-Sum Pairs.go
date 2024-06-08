package main

import "fmt"

func maxOperations(nums []int, k int) int {
	counts := make(map[int]int)
	result := 0

	for _, num := range nums {
		if counts[k-num] > 0 {
			result++
			counts[k-num]--
		} else {
			counts[num]++
		}
	}

	return result
}

func main() {
	fmt.Println(maxOperations([]int{1, 2, 3, 4}, 5))
	fmt.Println(maxOperations([]int{3, 1, 3, 4, 3}, 6))
	fmt.Println(maxOperations([]int{3, 1, 5, 1, 1, 1, 1, 1, 2, 2, 3, 2, 2}, 1))
}
