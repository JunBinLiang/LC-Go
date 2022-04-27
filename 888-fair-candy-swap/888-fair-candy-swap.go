func fairCandySwap(a []int, b []int) []int {
    res := make([]int, 2)
    var sum1, sum2 int = 0, 0
    f := make(map[int]bool)
    for i := 0; i < len(a); i++ {
        sum1 += a[i]
    }
    
    for i := 0; i < len(b); i++ {
        sum2 += b[i]
        f[b[i]] = true
    }
    
    for i := 0; i < len(a); i++ {
        dif := (sum2 - sum1) + 2 * a[i]
        if dif % 2 == 0 && f[dif / 2] {
            res[0] = a[i]
            res[1] = dif / 2
            break
        }
    }
    
    return res
}