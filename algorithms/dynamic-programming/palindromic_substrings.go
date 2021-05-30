package dynamicprogramming

/**
dp[i][j]标记s[i:j]的是否是回文串
状态转移方程：
    s[i:j]由单个字符组成；则 dp[i][j] = true
    s[i:j]由2个字符组成，且两个字符相同 s[i] = s[j]；则 dp[i][j] = true
    s[i:j]由超过2个字符组成，且首尾字符相同 s[i] = s[j]，且除去头尾剩余子串是一个回文串 dp[i+1][j-1] = true；则 dp[i][j] = tr
时间复杂度：O(n*n)
空间复杂度：O(n*n)
*/
func CountSubstrings(s string) int {
	n := len(s)
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
	}

	cnt := 0
	for j := 0; j < n; j++ {
		for i := 0; i <= j; i++ {
			if i == j {
				dp[i][j] = true
				cnt++
			} else if j-i == 1 && s[i] == s[j] {
				dp[i][j] = true
				cnt++
			} else if j-i > 1 && s[i] == s[j] && dp[i+1][j-1] {
				dp[i][j] = true
				cnt++
			}
		}
	}

	return cnt
}
