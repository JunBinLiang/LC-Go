func max(a int, b int) int {
    if a > b {
        return a
    }
    return b
}

func cal1(a []int) int {
    n := len(a)
    dp := make([][]int, n)
    for i := 0; i < n; i++ {
        dp[i] = make([]int, (n + 1) / 3 + 1)
        for j := 0; j < len(dp[0]); j++ {
            dp[i][j] = -1;
        }
    }
    
    dp[0][0] = 0
    dp[0][1] = a[0]
    
    for i := 1; i < n; i++ {
        dp[i][0] = 0
        for j := 1; j <= (n + 1) / 3; j++ {
            dp[i][j] = dp[i - 1][j]
            //take
            if i - 2 >= 0 {
                if dp[i - 2][j - 1] != -1 {
                    dp[i][j] = max(dp[i][j], a[i] + dp[i - 2][j - 1])
                }
            } else {
                dp[i][j] = max(dp[i][j], a[i])
            }
        }
    }
    
    return dp[n - 1][(n + 1) / 3]
}




func maxSizeSlices(a []int) int {
    n := len(a)
    b := make([]int, 0)
    c := make([]int, 0)
    for i := 0; i < n - 1; i ++ {
        b = append(b, a[i])
    }
    
    for i := 1; i < n; i ++ {
        c = append(c, a[i])
    } 
    
    return max(cal1(b), cal1(c));
}