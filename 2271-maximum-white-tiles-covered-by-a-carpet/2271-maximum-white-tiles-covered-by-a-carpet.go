func max(a int, b int) int {
    if a > b {
        return a
    }
    return b
}

func maximumWhiteTiles(a [][]int, l int) int {
    sort.SliceStable(a, func(i, j int) bool {
        return a[i][0] < a[j][0]
    })
    
    res := 0
    var sum, j int = 0, 0
    for i := 0; i < len(a); i++ {
        to := a[i][0] + l - 1
        for j < len(a) && to >= a[j][0] {
            sum += (a[j][1] - a[j][0] + 1)
            j++
        }
        if to < a[j - 1][1] {
            res = max(res, sum - (a[j - 1][1] - a[j - 1][0] + 1) + (to - a[j - 1][0] + 1))
        } else {
            res = max(res, sum)
        }
        sum -= (a[i][1] - a[i][0] + 1)
    }
    return res
}