func is(s string, l int, r int) bool {
    for l < r {
        if s[l] != s[r] {
            return false
        }
        l++
        r--
    }
    return true
}

func checkPartitioning(s string) bool {
    l := make([]bool, len(s))
    r := make([]bool, len(s))
    n := len(s)
    for i := 0; i < len(s); i ++ {
        l[i] = is(s, 0, i)
    }
    for i := len(s) - 1; i >= 0; i-- {
        r[i] = is(s, i, len(s) - 1)
    }
    
    for i := 0; i < len(s); i++ {
        var left, right int = i, i
        for left > 0 && right < n - 1 {
            if s[left] != s[right] {
                break
            }
            if l[left - 1] && r[right + 1] {
                return true
            } 
            left--
            right++
        }
    }
    
    for i := 0; i < len(s); i++ {
        var left, right int = i, i + 1
        for left > 0 && right < n - 1 {
            if s[left] != s[right] {
                break
            }
            if l[left - 1] && r[right + 1] {
                return true
            } 
            left--
            right++
        }
    }
    
    return false
}