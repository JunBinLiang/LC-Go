func min(a int, b int) int {
    if a < b {
        return a
    }
    return b
}

func minimumCardPickup(a []int) int {
    res := 1000000000
    f := make(map[int]int)
    for i := 0; i < len(a); i++ {
        if _, ok := f[a[i]]; ok {
            res = min(res, i - f[a[i]] + 1)
        }
        f[a[i]] = i
    }
    
    if res > len(a) {
        return -1
    }
    return res
}