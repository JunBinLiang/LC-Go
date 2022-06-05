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

func minMaxGame(a []int) int {
    for len(a) > 1 {
        res := make([]int, len(a) / 2)
        cnt := 0
        for i := 0; i < len(a); i += 2 {
            if cnt % 2 == 0 {
                res[i / 2] = min(a[i], a[i + 1])
            } else {
                res[i / 2] = max(a[i], a[i + 1])
            }
            cnt++
        }
        a = res
    }
    return a[0]
}