/*
 * @lc app=leetcode id=11 lang=golang
 *
 * [11] Container With Most Water
 */

// @lc code=start
func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func calculate(y0 int, y1 int, height []int, res *int) {
	minX := min(height[y0], height[y1])
	*res = max(minX*(y1-y0), *res)
}

func maxArea(height []int) int {
	res := 0
	for y0 := 0; y0 < len(height); y0++ {
		for y1 := 0; y1 < len(height); y1++ {
			calculate(y0, y1, height, &res)
		}
	}

	return res
}

// @lc code=end

