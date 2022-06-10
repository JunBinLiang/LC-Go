func min(a int, b int) int {
    if a < b {
        return a
    }
    return b
}

func dfs(state int, n int, k int, pre [][]int, dp []int) int {
    if state == (1 << n) - 1 {
        return 0
    }
    if dp[state] != -1 {
        return dp[state]
    }
    
    res := 1000000000
    
    a := make([]int, 0)
    for i := 0; i < n; i++ {
        if (state & (1 << i) ) == 0 { //not taken
            good := true
            for x := 0; x < len(pre[i]); x++ {
                if (state & (1 << pre[i][x])) == 0 {
                    good = false
                }
            }
            if good {
                a = append(a, i)
            }
        }
    }
        
    size := len(a)
    for i := 1; i < (1 << size); i++ {
        bitcnt := 0
        newtake := 0
        for j := 0; j < size; j++ {
            if (i & (1 << j)) > 0 {
                bitcnt++
                newtake |= (1 << a[j])
            }
        }
        
        if bitcnt <= k {
            res = min(res, 1 + dfs(state | newtake, n, k, pre, dp))
        }
    }
    
    dp[state] = res
    return dp[state]
}

func minNumberOfSemesters(n int, relations [][]int, k int) int {
    pre := make([][]int, n)
    for i := 0; i < len(pre); i++ {
        pre[i] = make([]int, 0)    
    }
    
    for i := 0; i < len(relations); i++ {
        var u, v int = relations[i][0] - 1, relations[i][1] - 1
        pre[v] = append(pre[v], u)
    }
    
    dp := make([]int, (1 << n))
    for i := 0; i < len(dp); i++ {
        dp[i] = -1
    }
    
      
    return dfs(0, n, k, pre, dp)
}