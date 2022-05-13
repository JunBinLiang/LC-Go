func get(a []int, l int, r int) int {
    if l == 0 {
        return a[r]
    }
    return a[r] ^ a[l - 1]
}

func xorQueries(a []int, q [][]int) []int {
    pre := make([]int, len(a))
    pre[0] = a[0]
    for i := 1; i < len(a); i++ {
        pre[i] = pre[i - 1] ^ a[i]
    }
    
    res := make([]int, len(q))
    for i := 0; i < len(q); i++ {
        res[i] = get(pre, q[i][0], q[i][1])
    }
    return res
}