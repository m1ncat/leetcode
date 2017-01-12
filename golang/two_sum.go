package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	numsSize := len(nums)

	var index1, index2 int
	for i := 0; i < numsSize; i++ {
		for j := i + 1; j < numsSize; j++ {
			if nums[i]+nums[j] == target {
				index1 = i
				index2 = j
				break
			}
		}
	}

	res := []int{index1, index2}
	return res
}

func main() {
	nums := []int{3, 2, 4}
	res := twoSum(nums, 6)

	fmt.Println(res)
}
