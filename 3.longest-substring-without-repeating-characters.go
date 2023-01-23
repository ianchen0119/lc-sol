package main

import "fmt"

/*
 * @lc app=leetcode id=3 lang=golang
 *
 * [3] Longest Substring Without Repeating Characters
 */

// @lc code=start

func max(m, n int) int {
	if m > n {
		return m
	}
	return n
}

func calculate(s string, start int, end int) (int, bool) {
	res := 0
	cur := 0
	filter := make([]int, 256)

	for start <= end {
		filter[int(s[start])] += 1
		if filter[int(s[start])] > 1 {
			cur = 1
			return res, true
		} else {
			cur += 1
		}
		start += 1
		res = max(res, cur)
	}
	return res, false
}

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	res := 1
	left := 0
	right := 0

	for right < len(s) {
		n := 0
		ok := false
		for n, ok = calculate(s, left, right); ok; {
			if left < right {
				left++
			}
			break
		}
		res = max(res, n)
		right++
	}
	return res
}

// @lc code=end

func main() {
	s := "abcabcbb"
	fmt.Println(lengthOfLongestSubstring(s))
}
