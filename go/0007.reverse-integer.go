import "math"

func reverse(x int) int {
	var min = math.MinInt32
	var max = math.MaxInt32
	var out int
	for x != 0 {
		digit := x % 10
		out = out*10 + digit
		if out > max || out < min {
			return 0
		}
		x = x / 10
	}
	return int(out)
}
