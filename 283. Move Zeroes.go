package main

import "fmt"

func moveZeroes(nums []int) {
	lastNonZeroIndex := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[lastNonZeroIndex], nums[i] = nums[i], nums[lastNonZeroIndex]
			lastNonZeroIndex++
		}
	}
}

func main() {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
	fmt.Println(nums)

	nums = []int{0}
	moveZeroes(nums)
	fmt.Println(nums)
}
