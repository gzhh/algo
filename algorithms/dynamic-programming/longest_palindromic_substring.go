package dynamicprogramming

func LongestPalindrome(s string) string {
	n := len(s)
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
	}

	ans := ""
	for k := 0; k < n; k++ {
		for i := 0; i+k < n; i++ {
			j := i + k
			if k == 0 {
				dp[i][j] = true
			} else if k == 1 && s[i] == s[i+1] {
				dp[i][j] = true
			} else {
				if s[i] == s[j] && dp[i+1][j-1] {
					dp[i][j] = true
				}
			}
			if dp[i][j] && k+1 > len(ans) {
				ans = s[i : i+k+1]
			}
		}
	}

	return ans
}
