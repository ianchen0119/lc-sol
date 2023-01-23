package main

import "fmt"

/*
 * @lc app=leetcode id=15 lang=golang
 *
 * [15] 3Sum
 */

// @lc code=start
func sort(nums []int) {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] > nums[j] {
				temp := nums[i]
				nums[i] = nums[j]
				nums[j] = temp
			}
		}
	}
}

func isDuplicated(nums []int, filter map[string]bool) bool {
	if len(nums) != 3 {
		return false
	}
	sort(nums)
	var key string
	key = fmt.Sprintf("%d%d%d", nums[0], nums[1], nums[2])

	if filter[key] {
		return true
	} else {
		filter[key] = true
		return false
	}
}

func threeSum(nums []int) [][]int {
	var filter map[string]bool
	filter = make(map[string]bool)
	var res [][]int
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			for k := j + 1; k < len(nums); k++ {
				if nums[i]+nums[j]+nums[k] == 0 {
					slice := []int{nums[i], nums[j], nums[k]}
					if !isDuplicated(slice, filter) {
						res = append(res, slice)
					}
				}
			}
		}
	}
	return res
}

// @lc code=end

func main() {
	test := []int{-1, 0, 1}
	fmt.Println(threeSum(test))
	// sort(test)
	// fmt.Println(test)
}
