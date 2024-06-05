package main

import "fmt"

func isSubsequence(s string, t string) bool {
	if s == "" {
		return true
	} else if t == "" {
		return false
	}
	indexFound := 0
	runes := []rune(s)
	for _, char := range t {
		if char == runes[indexFound] {
			indexFound += 1
		}
		if indexFound == len(s) {
			return true
		}
	}
	return false

}

func main() {
	fmt.Println(isSubsequence("abc", "ahbgdc"))
	fmt.Println(isSubsequence("axc", "ahbgdc"))
}
