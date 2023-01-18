/*
 * @lc app=leetcode id=5 lang=golang
 *
 * [5] Longest Palindromic Substring
 */

// @lc code=start
func isPalindrome(s string, i int, j int) bool {
	if i >= j {
		return false
	}
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func longestPalindrome(s string) string {
	res := s[:1]
	length := 1
	for i := 0; i < len(s); i++ {
		for j := i; j < len(s); j++ {
			if isPalindrome(s, i, j) {
				if len(s[i:j+1]) > length {
					res = s[i : j+1]
					length = len(s[i : j+1])
				}
			}
		}
	}
	return res
}

// @lc code=end

