func most(a []int, k int) int {
    f := make(map[int]int)
    var j, res int = 0, 0
    for i := 0; i < len(a); i++ {
        f[a[i]]++
        for len(f) > k {
            f[a[j]]--
            if f[a[j]] == 0 {
                delete(f, a[j])
            }
            j++
        }
        res += (i - j + 1)
    }
    return res
}

func subarraysWithKDistinct(a []int, k int) int {
    return most(a, k) - most(a, k - 1)    
}