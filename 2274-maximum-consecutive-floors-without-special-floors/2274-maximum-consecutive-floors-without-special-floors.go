func maxConsecutive(l int, r int, a []int) int {
    a = append(a, l - 1)
    a = append(a, r + 1)
    sort.Ints(a)
    res := 0
    for i := 1; i < len(a); i++ {
        dif := a[i] - a[i - 1] - 1
        if dif > res {
            res =dif
        }
    }
    return res
}