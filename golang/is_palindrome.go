package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func isPalindrome(x int) bool {
	arr := []byte(strconv.Itoa(x))
	length := len(arr)

	var firstPart []byte
	var lastPart []byte

	if length%2 == 0 {
		for i := 0; i < length/2; i++ {
			firstPart = append(firstPart, arr[i])
		}

		for j := length - 1; j >= length/2; j-- {
			lastPart = append(lastPart, arr[j])
		}
	} else {
		for i := 0; i < length/2; i++ {
			firstPart = append(firstPart, arr[i])
		}

		for j := length - 1; j > length/2; j-- {
			lastPart = append(lastPart, arr[j])
		}
	}

	fmt.Println(firstPart)
	fmt.Println(lastPart)

	var flag bool
	if reflect.DeepEqual(firstPart, lastPart) {
		flag = true
	} else {
		flag = false
	}

	return flag
}

func main() {
	test := 12321
	flag := isPalindrome(test)
	fmt.Println(flag)
}
