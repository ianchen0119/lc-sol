package main

import (
	"fmt"
	"sort"
)

/*
 * @lc app=leetcode id=15 lang=golang
 *
 * [15] 3Sum
 */

// @lc code=start

func threeSum(nums []int) [][]int {
	var result [][]int
	if len(nums) < 3 {
		return result
	}

	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		if i != 0 && nums[i-1] == nums[i] {
			continue
		}

		for k, j := i+1, len(nums)-1; k < j; {
			n := nums[i] + nums[k] + nums[j]
			if n == 0 {
				result = append(result, []int{nums[i], nums[k], nums[j]})
				l := k
				for l < j && nums[l] == nums[k] {
					l++
				}
				k = l
			} else if n > 0 {
				j--
			} else {
				k++
			}
		}
	}
	return result
}

// @lc code=end

func main() {
	test := []int{-1, 0, 1}
	fmt.Println(threeSum(test))
	// sort(test)
	// fmt.Println(test)
}
