package main

import (
	"fmt"
)

func removeDuplicates(nums []int) int {

	k := 1

	for i := range nums {
		if nums[0] != nums[i] && nums[k-1] != nums[i] {
			nums[k] = nums[i]
			k++
		}

	}
	fmt.Println(nums)
	fmt.Println(k)
	return k

}

func main() {

	removeDuplicates([]int{-1, 0, 0, 0, 0, 3, 3})
}
