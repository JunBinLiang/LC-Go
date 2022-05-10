var res [][]int
var K int
func dfs(start int, sum int, cur []int) {
    if sum < 0 {
        return
    }
    
    if len(cur) == K {
        if sum == 0 {
            a := make([]int, K)
            for i := 0; i < K; i++ {
                a[i] = cur[i]
            }
            res = append(res, a)
        } 
        return
    }
    
    for i := start; i <= 9; i++ {
        cur = append(cur, i)
        dfs(i + 1, sum - i, cur)
        cur = cur[ : len(cur) - 1]
    }
    
}

func combinationSum3(k int, n int) [][]int {
    K = k
    res = make([][]int, 0)
    cur := make([]int, 0)
    dfs(1, n, cur)    
    return res
}