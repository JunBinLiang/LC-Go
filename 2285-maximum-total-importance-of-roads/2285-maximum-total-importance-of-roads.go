func maximumImportance(n int, roads [][]int) int64 {
    cnt := make([]int, n)
    for i := 0; i < len(roads); i++ {
        cnt[roads[i][0]]++
        cnt[roads[i][1]]++
    }
    
    sort.Ints(cnt)
    res := int64(0)
    for i := 0; i < len(cnt); i++ {
        res += int64(i + 1) * int64(cnt[i])
    }
    return res
}