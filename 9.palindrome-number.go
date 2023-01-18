/*
 * @lc app=leetcode id=9 lang=golang
 *
 * [9] Palindrome Number
 */

// @lc code=start
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	s := strconv.Itoa(x)
	cond := len(s)
	for i := 0; i < cond; i++ {
		if s[i] != s[cond-i-1] {
			return false
		}
		if i > cond/2 {
			return true
		}
	}
	return true
}

// @lc code=end

