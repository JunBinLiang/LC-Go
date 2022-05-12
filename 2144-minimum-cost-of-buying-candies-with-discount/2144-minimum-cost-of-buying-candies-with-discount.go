func minimumCost(a []int) int {
    sort.Ints(a)
    res := 0
    for i := len(a) - 1; i >= 0; i-- {
        res += a[i]
        if i - 1 >= 0 {
            res += a[i - 1]
        }
        i -= 2
    }
    return res
}