package main

import (
	"fmt"
)

func length_of_the_string(s string) int {
	length := len(s)
	if length == 0 {
		return 0
	}
	return length
}

func number_of_each_letter(s string) map[string]int {
	number := make(map[string]int)
	for _, v := range s {
		number[string(v)]++
	}
	return number
}

func longest_string(s string) int {
	number := number_of_each_letter(s)
	number_of_logest_substring := 0
	for _, v := range number {
		if v%2 == 0 {
			number_of_logest_substring += v
		} else {
			number_of_logest_substring += v - 1
		}
	}
	return number_of_logest_substring
}

func main() {
	a := "abcabcbbsss"
	fmt.Println("字符串的长度是:", length_of_the_string(a))
	fmt.Println(number_of_each_letter(a))
	fmt.Println("所以最长回文串的长度是:", longest_string(a))
}
