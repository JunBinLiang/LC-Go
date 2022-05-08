func countTexts(s string) int {
    mod := 1000000007
    a := []int{0, 0, 3, 3, 3, 3, 3, 4, 3, 4}
    dp := make([]int, len(s))
    for i := 0; i < len(s); i++ {
        t := a[s[i] - '0']
        for j := 0; j < t; j++ {
            if i - j >= 0 && s[i - j] == s[i] {
                if i - j - 1 < 0 {
                    dp[i] ++
                } else {
                    dp[i] += dp[i - j - 1]   
                }
                dp[i] %= mod
            } else {
                break
            }
        }
    }
    return dp[len(s) - 1]
    
}