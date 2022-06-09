func max(a int, b int) int {
    if a > b {
        return a
    }
    return b
}

func maxAscendingSum(a []int) int {
    n := len(a)
    res := 0
    for i := 0; i < n; i++ {
        sum := a[i]
        for j := i + 1; j < n; j++ {
            if a[j] > a[j - 1] {
                sum += a[j]
            } else {
                break
            }
        }
        res = max(res, sum)
    }
    return res
}