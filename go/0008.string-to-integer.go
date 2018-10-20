import "math"

func myAtoi(str string) int {
	n := len(str)
	if n == 0 {
		return 0
	}

	i := 0

	// Skip white spaces
	for i < n && str[i] == ' ' {
		i++
	}

	// Read sign
	sign := 1
	if i < n {
		if str[i] == '-' {
			sign = -1
			i++
		} else if str[i] == '+' {
			sign = 1
			i++
		}
	}

	ret := 0
	for i < n && str[i] >= '0' && str[i] <= '9' {
		ret = ret*10 + int(str[i]-'0')
		if ret > math.MaxInt32 {
			break
		} else {
			i++
		}
	}
	ret = ret * sign

	// Check overflow
	if ret < math.MinInt32 {
		return math.MinInt32
	}
	if ret > math.MaxInt32 {
		return math.MaxInt32
	}
	return ret
}
