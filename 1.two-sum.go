/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 */
package main

import "fmt"

// @lc code=start
func twoSum(nums []int, target int) []int {
	// var cache map[int]int
	cache := make(map[int]int, len(nums))
	for key, value := range nums {
		cache[value] = key
	}
	for key, value := range nums {
		if val, ok := cache[target-value]; ok && val != key {
			return []int{key, val}
		}
	}
	return []int{}
}

// @lc code=end
func main() {
	// twoSum([]int{3, 2, 4}, 6)
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
}
