func min(a int, b int) int {
    if a < b {
        return a
    }
    return b
}

func minimumTotal(a [][]int) int {
    for i := len(a) - 2; i >= 0; i-- {
        for j := 0; j < len(a[i]); j++ {
            a[i][j] += min(a[i + 1][j], a[i + 1][j + 1])
        }
    }
    return a[0][0]
}