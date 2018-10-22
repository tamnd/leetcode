func maxArea(height []int) int {
	i, j := 0, len(height)-1
	ret := 0

	for i < j {
		a, b := height[i], height[j]
		h := min(a, b)
		area := h * (j - i)
		if ret < area {
			ret = area
		}

		if a < b {
			i++
		} else {
			j--
		}
	}
	return ret
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
