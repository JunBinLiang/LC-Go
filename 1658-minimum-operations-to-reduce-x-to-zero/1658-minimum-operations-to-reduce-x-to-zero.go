func min(a int, b int) int {
    if a < b {
        return a
    }
    return b
}

func minOperations(a []int, x int) int {
    res := 1000000000
    sum := 0
    
    right := make([]int, len(a))
    for i := len(a) - 1; i >= 0; i-- {
        sum += a[i]
        right[i] = sum
        if sum == x {
            res = min(res, len(a) - i)
        }
    }
    
    sum = 0
    for i := 0; i < len(a); i++ {
        sum += a[i]
        if sum == x {
            res = min(res, i + 1)
        }
        var l, r int = i + 1, len(a) - 1
        for l <= r {
            mid := l + (r - l) / 2
            if right[mid] + sum == x {
                res = min(res, i + 1 + (len(a) - mid))
                break
            } else if right[mid] + sum > x {
                l = mid + 1
            } else {
                r = mid - 1
            }
        }
    }
    
    if res == 1000000000 {
        return -1
    }
    return res
}