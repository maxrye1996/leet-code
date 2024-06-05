package main

import "fmt"

func mergeAlternately(word1 string, word2 string) string {
    length1 := len(word1)
    length2 := len(word2)
    merged := make([]byte, 0, length1+length2)
    
    i, j := 0, 0
    for i < length1 || j < length2 {
        if i < length1 {
            merged = append(merged, word1[i])
            i++
        }
        if j < length2 {
            merged = append(merged, word2[j])
            j++
        }
    }

    return string(merged)
}

func main() {
    word1 := "ab"
    word2 := "cdefghijklmnopqrstuvwxyz"
    result := mergeAlternately(word1, word2)
    fmt.Println(result) // Output: "adbecfgh"
}