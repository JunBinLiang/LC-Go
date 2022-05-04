func maxOperations(a []int, k int) int {
    f := make(map[int]int)
    res := 0
    for i := 0; i < len(a); i++ {
        f[a[i]]++
    }
    
    for i := 0; i < len(a); i++ {
        if f[a[i]] > 0 {
            need := k - a[i]
            f[a[i]]--
            if f[need] > 0 {
                res++
                f[need]--
            }
        }
    }
    return res
}