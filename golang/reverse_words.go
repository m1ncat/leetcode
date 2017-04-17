package main

import (
	"fmt"
	"strings"
)

func reverseString(s string) string {
	runes := []rune(s)
	for from, to := 0, len(runes)-1; from < to; from, to = from+1, to-1 {
		runes[from], runes[to] = runes[to], runes[from]
	}
	return string(runes)
}

func reverseWords(s string) string {
	arr := strings.Split(s, " ")
	var reverseArr []string
	for _, ele := range arr {
		reverseStr := reverseString(ele)
		reverseArr = append(reverseArr, reverseStr)
	}

	var res string
	for _, ele := range reverseArr {
		res += ele + " "
	}
	index := strings.LastIndex(res, " ")
	ret := res[:index]
	return ret
}

func main() {
	testStr := "Let's take LeetCode contest"
	res := reverseWords(testStr)
	fmt.Println(res)
}
