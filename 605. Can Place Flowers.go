package main

import (
	"fmt"
)

func canPlaceFlowers(flowerbed []int, n int) bool {
	numberPlaced := 0
	zeroCount := 0
	startPlacing := false
	for i, value := range flowerbed {
		if i == 0 && value == 0 {
			zeroCount += 2
		} else if value == 0 {
			zeroCount += 1
		} else {
			zeroCount = 0
			startPlacing = false
		}
		if zeroCount == 3 {
			startPlacing = true
			numberPlaced += 1
		} else if zeroCount%2 == 1 && startPlacing {
			numberPlaced += 1
		} else if i+1 == len(flowerbed) && zeroCount > 1 {
			numberPlaced += 1
		}

	}
	if numberPlaced >= n {
		return true
	} else {
		return false
	}
}

func main() {
	fmt.Println(canPlaceFlowers([]int{1, 0, 0, 0, 1}, 1))
	fmt.Println(canPlaceFlowers([]int{1, 0, 0, 0, 1}, 2))
	fmt.Println(canPlaceFlowers([]int{1, 0, 0, 0, 0, 0, 0, 0, 1}, 3))
	fmt.Println(canPlaceFlowers([]int{0, 0, 1, 0, 1}, 1))
	fmt.Println(canPlaceFlowers([]int{1, 0, 1, 0, 0}, 1))
	fmt.Println(canPlaceFlowers([]int{1, 0}, 1))
}
