func max(a int, b int) int {
    if a > b {
        return a
    }
    return b
}

func largestCombination(a []int) int {
    res := 0
    for i := 0; i < 25; i++ {
        cnt := 0
        for _, v := range a {
            if (v & (1 << i)) > 0 {
                cnt++
            }
        }
        res = max(res, cnt)
    }
    return res
}