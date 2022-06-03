type NumMatrix struct {
    DP [][]int
}


func Constructor(grid [][]int) NumMatrix {
    var n, m int = len(grid), len(grid[0])
    dp := make([][]int, n + 1)
    for i := 0; i <= n; i++ {
        dp[i] = make([]int, m + 1)
    }
    
    for i := 1; i <= n; i++ {
        for j := 1; j <= m; j++ {
            dp[i][j] = grid[i - 1][j - 1] + dp[i - 1][j] + dp[i][j - 1] - dp[i - 1][j - 1]
        }
    }
    
    return NumMatrix{dp}
}


func (this *NumMatrix) SumRegion(r1 int, c1 int, r2 int, c2 int) int {
    dp := this.DP
    return dp[r2 + 1][c2 + 1] + dp[r1][c1] - dp[r1][c2 + 1] - dp[r2 + 1][c1];
}


/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */