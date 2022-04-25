func mostCompetitive(a []int, k int) []int {
    n := len(a)
    res := make([]int, 0)
    for i := 0; i < len(a); i++ {
        for len(res) > 0 && a[i] < res[len(res) - 1] && len(res) + (n - i - 1) >= k {
            res = res[ : len(res) - 1]
        }
        res = append(res, a[i])
    }
    
    for len(res) > k {
         res = res[ : len(res) - 1]
    }
    
    return res
}