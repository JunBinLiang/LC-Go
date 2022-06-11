func get(a []int64, l int, r int) int64 {
    if l == 0 {
        return a[r]
    }
    return a[r] - a[l - 1]
}

func countSubarrays(a []int, k int64) int64 {
    res := int64(0)
    
    pre := make([]int64, len(a))
    pre[0] = int64(a[0])
    for i := 1; i < len(pre); i++ {
        pre[i] = pre[i - 1] + int64(a[i])
    }
    
    for i := 0; i < len(a); i++ {
        var l, r int = i, len(a) - 1
        pos := -1
        
        for l <= r {
            mid := l + (r - l) / 2
            size := int64(mid - i + 1)
            sum := get(pre, i, mid)
            if sum * size < k {
                pos = mid
                l = mid + 1
            } else {
                r = mid - 1
            }
        }
        
        if pos != -1 {
            res += int64((pos - i + 1))
        }
    }
    return res
}