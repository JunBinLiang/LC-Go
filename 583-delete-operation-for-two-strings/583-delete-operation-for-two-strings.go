func max(a int, b int) int {
    if a > b {
        return a
    }
    return b
}

func minDistance(s string, t string) int {
    var n, m int = len(s), len(t)
    dp := make([][]int, n + 1)
    for i := 0; i < len(dp); i++ {
        dp[i] = make([]int, m + 1)
    }
    
    for i := 1; i <= n; i++ {
        for j := 1; j <= m; j++ {
            if s[i - 1] == t[j - 1] {
                dp[i][j] = 1 + dp[i - 1][j - 1]
            } else {
                dp[i][j] = max(dp[i - 1][j], dp[i][j - 1])
            }
        }
    }
    return n + m - dp[n][m] * 2
}