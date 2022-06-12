func max(a int, b int) int {
    if a > b {
        return a
    }
    return b
}

func maximumUniqueSubarray(a []int) int {
    res := 0
    sum := 0
    j := 0
    f := make(map[int]bool)
    for i := 0; i < len(a); i++ {
        for f[a[i]] == true {
            f[a[j]] = false
            sum -= a[j]
            j++
        }
        
        sum += a[i]
        f[a[i]] = true
        res = max(res, sum)
    }
    
    return res
}