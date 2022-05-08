var dp[101][101][110]bool

func hasValidPath(a [][]byte) bool {
    if a[0][0] == ')' {
        return false
    }
    
    var n, m int = len(a), len(a[0])
    for i := 0; i < n; i++ {
        for j := 0; j < m; j++ {
            for k := 0; k < 101; k++ {
                dp[i][j][k] = false
            }
        }
    }
    
    dp[0][0][1] = true
    for i := 0; i < n; i++ {
        for j := 0; j < m; j++ {
            if i - 1 >= 0 {
                if a[i][j] == '(' {
                    for open := 1; open < 105; open ++ {
                        dp[i][j][open] = dp[i][j][open] || dp[i - 1][j][open - 1]
                    }
                } else {
                    for open := 0; open < 105; open ++ {
                        dp[i][j][open] = dp[i][j][open] || dp[i - 1][j][open + 1]
                    }
                }
            }
            if j - 1 >= 0 {
                if a[i][j] == '(' {
                    for open := 1; open < 105; open ++ {
                        dp[i][j][open] = dp[i][j][open] || dp[i][j - 1][open - 1]
                    }
                } else {
                    for open := 0; open < 105; open ++ {
                        dp[i][j][open] = dp[i][j][open] || dp[i][j - 1][open + 1]
                    }
                }
            }
        }
    }
    return dp[n - 1][m - 1][0]
}