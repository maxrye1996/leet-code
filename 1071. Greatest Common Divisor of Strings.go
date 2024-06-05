package main

import "fmt"

func gcdOfStrings(str1 string, str2 string) string {

	if str1+str2 != str2+str1 {
		return ""
	}

	tmpOne := len(str1)
	tmpTwo := len(str2)
	for tmpTwo != 0 {
		tmpOne, tmpTwo = tmpTwo, tmpOne%tmpTwo
	}

	return str1[:tmpOne]
}

func main() {
	testOne := "ababab"
	testTwo := "abab"
	fmt.Println(gcdOfStrings(testOne, testTwo))
}
