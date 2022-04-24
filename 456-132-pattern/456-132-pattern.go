func reverse(a []int) {
    var l, r int = 0, len(a) - 1
    for l < r {
        t := a[l]
        a[l] = a[r]
        a[r] = t
        l++
        r--
    }
}

func min(a int, b int) int {
    if a < b {
        return a
    }
    return b
}

func max(a int, b int) int {
    if a > b {
        return a
    }
    return b
}


func find132pattern(a []int) bool {
    reverse(a)
    r := make([]int, len(a))
    mn := 1000000000
    for i := len(a) - 1; i >= 0; i-- {
        mn = min(mn, a[i])
        r[i] = mn
    }
    
  
    sta := make([]int, 0)
    for i := 0; i < len(a) - 1; i++ {
        mx := -1000000000
        for len(sta) > 0 && a[i] > sta[len(sta) - 1] {
            last := sta[len(sta) - 1]
            sta = sta[ : len(sta) - 1]
            mx = max(mx, last)
        }
        sta = append(sta, a[i])
        
        if r[i + 1] < mx {
            return true
        }
    }
    return false
}
