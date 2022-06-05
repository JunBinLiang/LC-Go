func partitionArray(a []int, k int) int {
    sort.Ints(a)
    res := 0
    
    first := a[0]
    for i := 0; i < len(a); i++ {
        if a[i] - first > k {
            res++
            first = a[i]
        }
    }
    
    return res + 1
}