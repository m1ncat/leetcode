package main

import (
	"fmt"
	"strconv"
)

func reverse(x int) int {
	var value int
	var ret int
	if x >= 0 {
		value = x
	} else {
		value = -x
	}

	str := strconv.Itoa(value)
	var arr []byte
	for i := len(str) - 1; i >= 0; i-- {
		arr = append(arr, str[i])
	}
	ret, _ = strconv.Atoi(string(arr))

	if x >= 0 {
		ret = ret
	} else {
		ret = -ret
	}

	if ret >= 2147483647 || ret <= -2147483648 {
		ret = 0
	}

	return ret
}

func main() {
	res := reverse(-2147483648)
	fmt.Println(res)
}
