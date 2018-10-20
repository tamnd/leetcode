func lengthOfLongestSubstring(s string) int {
	ret := 0
	var arr [256]int
	start := 0
	for i, ch := range s {
		if arr[ch] > start {
			start = arr[ch]
		}
		if i-start+1 > ret {
			ret = i - start + 1
		}
		arr[ch] = i + 1
	}
	return ret
}
