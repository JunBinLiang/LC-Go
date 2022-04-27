func max(a int, b int) int {
    if a > b {
        return a
    }
    return b
}

func findLengthOfLCIS(a []int) int {
    var res, cnt int = 1, 1
    for i := 1; i < len(a); i++ {
        if a[i] > a[i - 1] {
            cnt++
        } else {
            cnt = 1
        }
        res = max(res, cnt)
    }
    return res
}