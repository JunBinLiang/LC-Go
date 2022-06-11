func successfulPairs(a []int, b []int, success int64) []int {
    sort.Ints(b)
    res := make([]int, 0)
    for i := 0; i < len(a); i++ {
        var l, r int = 0, len(b) - 1
        pos := -1
        for l <= r {
            mid := l + (r - l) / 2
            x := int64(a[i])
            y := int64(b[mid])
            if x * y < success {
                pos = mid
                l = mid + 1
            } else {
                r = mid - 1
            }
            
        }
        res = append(res, (len(b)) - (pos + 1))
    }
    return res
}