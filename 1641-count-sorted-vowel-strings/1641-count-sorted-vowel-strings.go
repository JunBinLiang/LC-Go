func countVowelStrings(n int) int {
    dp := make([][]int, n)
    for i := 0; i < n; i++ {
        dp[i] = make([]int, 5)
    }
    
    for i := 0; i < 5; i++ {
        dp[0][i] = 1
    }
    
    for i := 1; i < n; i++ {
        sum := 0
        for j := 0; j < 5; j++ {
            sum += dp[i - 1][j]
            dp[i][j] = sum
        }
    }
    
    
    res := 0
    for i := 0; i < 5; i++ {
        res += dp[n - 1][i]
    }
    return res
}