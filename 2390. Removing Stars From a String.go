package main

import "fmt"

func removeStars(s string) string {
	result := []rune{}

	for _, char := range s {
		if char == '*' {
			if len(result) > 0 {
				result = result[:len(result)-1]
			}
		} else {
			result = append(result, char)
		}
	}

	return string(result)
}

func main() {
	fmt.Println(removeStars("leet**cod*e"))
	fmt.Println(removeStars("erase*****"))
}
