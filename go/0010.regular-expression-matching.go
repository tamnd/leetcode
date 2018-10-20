func isMatch(s string, p string) bool {
	lens := len(s)
	lenp := len(p)

	dp := make([][]bool, lens+1)
	for i := range dp {
		dp[i] = make([]bool, lenp+1)
	}

	dp[0][0] = true
	for j := 1; j < lenp && dp[0][j-1]; j += 2 {
		if p[j] == '*' {
			dp[0][j+1] = true
		}
	}

	for i := 0; i < lens; i++ {
		for j := 0; j < lenp; j++ {
			if p[j] == '.' || p[j] == s[i] {
				dp[i+1][j+1] = dp[i][j]
			} else if p[j] == '*' {
				if p[j-1] != s[i] && p[j-1] != '.' {
					dp[i+1][j+1] = dp[i+1][j-1]
				} else {
					dp[i+1][j+1] = dp[i+1][j-1] || dp[i+1][j] || dp[i][j+1]
				}
			}
		}
	}

	return dp[lens][lenp]
}
