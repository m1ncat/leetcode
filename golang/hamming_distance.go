package main

import (
	"fmt"
	"strconv"
	"strings"
)

func hammingDistance(x int, y int) int {
	xorNum := int64(x ^ y)
	s := strconv.FormatInt(xorNum, 2)
	res := strings.Count(s, "1")
	return res
}

func main() {
	x := 1
	y := 4
	res := hammingDistance(x, y)
	fmt.Println(res)
}
