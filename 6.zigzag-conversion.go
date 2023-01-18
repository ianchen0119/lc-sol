/*
 * @lc app=leetcode id=6 lang=golang
 *
 * [6] Zigzag Conversion
 */

// @lc code=start

func zigZag(s string, numRows int, ss []string) []string {
	i := 0
	j := 0

	for i < len(s) {
		j = 0
		for j < numRows {
			if i < len(s) {
				ss[j] += string(s[i])
			} else {
				return ss
			}
			i++
			j++
			if j == numRows-1 {
				for j > 0 {
					if i < len(s) {
						ss[j] += string(s[i])
					} else {
						return ss
					}

					i++
					j--
				}
			}
		}
	}
	return ss
}

func convert(s string, numRows int) string {
	ss := make([]string, numRows)
	zigZag(s, numRows, ss)
	res := ""
	for i := 0; i < len(ss); i++ {
		res += ss[i]
	}
	return res
}

// @lc code=end

