package main

import "fmt"

/*
 * @lc app=leetcode id=4 lang=golang
 *
 * [4] Median of Two Sorted Arrays
 */

// @lc code=start

// Input: nums1 = [1,3], nums2 = [2]
// Output: 2.00000
// Explanation: merged array = [1,2,3] and median is 2.

func findTarget(nums1 []int, nums2 []int, n int) (int, int) {
	i := 0
	j := 0
	cur := 0
	prev := 0
	for i < len(nums1) || j < len(nums2) {
		if i < len(nums1) && j < len(nums2) {
			if nums1[i] >= nums2[j] {
				prev = cur
				cur = nums2[j]
				j++
			} else {
				prev = cur
				cur = nums1[i]
				i++
			}
		} else if i < len(nums1) {
			prev = cur
			cur = nums1[i]
			i++
		} else if j < len(nums2) {
			prev = cur
			cur = nums2[j]
			j++
		}
		if i+j == n {
			return prev, cur
		}
	}
	return prev, cur
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	n := len(nums1) + len(nums2)
	if n == 0 {
		return 0
	}
	isOdd := false
	if n&1 > 0 {
		isOdd = true
	}
	target := n/2 + 1
	if isOdd {
		_, f2 := findTarget(nums1, nums2, target)
		return float64(f2)
	} else {
		f1, f2 := findTarget(nums1, nums2, target)
		return (float64(f1) + float64(f2)) / 2
	}
}

// @lc code=end
func main() {
	nums1 := []int{}
	nums2 := []int{2}
	fmt.Println(findMedianSortedArrays(nums1, nums2))
}
