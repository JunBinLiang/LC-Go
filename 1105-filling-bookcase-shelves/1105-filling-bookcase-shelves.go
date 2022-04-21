func max(a int, b int) int {
    if a > b {
        return a
    }
    return b
}

func min(a int, b int) int {
    if a < b {
        return a
    }
    return b
}

func minHeightShelves(a [][]int, w int) int {
    n := len(a)
    dp := make([]int, n)
    for i := 0; i < n; i++ {
        dp[i] = 1000000000 //init max
    }
    
    for i := 0; i < n; i++ {
        var sum, h int = 0, 0
        for j := i; j >= 0; j-- {
            sum += a[j][0]
            if sum > w {
                break
            }
            h = max(h, a[j][1])
            if j - 1 >= 0 {
                dp[i] = min(dp[i], h + dp[j - 1])
            } else {
                dp[i] = h
            }
        }
    }
    return dp[n - 1]
}