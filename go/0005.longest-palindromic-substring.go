func longestPalindrome(s string) string {
	length := len(s)
	if length < 2 {
		return s
	}
	start, maxLen := 0, 1

	for i := 0; i < length; {
		if length <= maxLen/2+i {
			break
		}
		left, right := i, i
		for right < length-1 && s[right+1] == s[right] {
			right++
		}
		i = right + 1
		for right < length-1 && left > 0 && s[right+1] == s[left-1] {
			right++
			left--
		}
		newLen := right + 1 - left
		if newLen > maxLen {
			start = left
			maxLen = newLen
		}
	}
	return s[start : start+maxLen]
}
