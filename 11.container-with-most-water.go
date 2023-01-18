/*
 * @lc app=leetcode id=11 lang=golang
 *
 * [11] Container With Most Water
 */

// @lc code=start
func min(x, y int) (int, bool) {
	if x > y {
		return y, false
	}
	return x, true
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func maxArea(height []int) int {
	res := 0
	y0 := 0
	y1 := len(height) - 1
	for y0 <= y1 {
		minX, cond := min(height[y0], height[y1])
		res = max(minX*(y1-y0), res)
		// if height[y0] is smaller than height[y1]
		if cond {
			y0 += 1
		} else {
			y1 -= 1
		}
	}

	return res
}

// @lc code=end

