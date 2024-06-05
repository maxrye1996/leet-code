package main

import (
	"fmt"
)

func maxArea(height []int) int {
	l, r := 0, len(height)-1
	biggestArea := 0
	area := 0

	for l < r {
		if height[l] < height[r] {
			area = (r - l) * height[l]
		} else {
			area = (r - l) * height[r]
		}

		if area > biggestArea {
			biggestArea = area
		}
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}

	return biggestArea
}

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
	fmt.Println(maxArea([]int{1, 1}))
	fmt.Println(maxArea([]int{2, 3, 4, 5, 18, 17, 6}))
}
