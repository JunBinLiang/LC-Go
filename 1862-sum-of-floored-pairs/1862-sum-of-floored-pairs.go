func get(a []int, l int, r int) int {
    if l > r {
        return 0
    }
    
    if l == 0 {
        return a[r]
    }
    
    return a[r] - a[l - 1]
}

func min(a int, b int) int {
    if a < b {
        return a
    }
    return b
}

func sumOfFlooredPairs(a []int) int {
    mod := 1000000007
    sort.Ints(a)
    res := 0
    
    cnt := make([]int, a[len(a) - 1] + 1)
    for i := 0; i < len(a); i++ {
        cnt[a[i]]++
    }
    
    for i := 1; i < len(cnt); i++ {
        cnt[i] += cnt[i - 1]
    }
    
    f := make(map[int]int)
    
    for i := 0; i < len(a); i++ {
        num := a[i]
        if _, ok := f[num]; ok {
            res += f[num]
            res %= mod
            continue
        }
        
        total := 0
        for j := num; j <= a[len(a) - 1]; j += num {
            mul := j / num
            sum := get(cnt, j, min(j + num - 1, a[len(a) - 1]))
            total += (mul * sum)
            total %= mod
        }
        f[num] = total
        res += total
        res %= mod
    }
    
    return res
}