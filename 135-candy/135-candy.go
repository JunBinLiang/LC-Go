func max(a int, b int) int {
    if a > b {
        return a
    }
    return b
}

func candy(a []int) int {
    n := len(a)
    
    graph := make([][]int, n)
    in := make([]int, n)
    for i := 0; i < n; i++ {
        graph[i] = make([]int, 0)
    }
    
    for i := 0; i < n; i++ {
        if i + 1 < n && a[i + 1] > a[i] {
            graph[i] = append(graph[i], i + 1)
            in[i + 1]++
        } 
        
        if i - 1 >= 0 && a[i - 1] > a[i] {
            graph[i] = append(graph[i], i - 1)
            in[i - 1]++
        }
    }
    
    // do topological sort
    
    q := make([]int, 0)
    dis := make([]int, n)
    for i := 0; i < n; i++ {
        if in[i] == 0 {
            q = append(q, i)
            dis[i] = 1
        }
    }
    
    
    for len(q) > 0 {
        top := q[0]
        q = q[1 :]
        for i := 0; i < len(graph[top]); i++ {
            nxt := graph[top][i]
            in[nxt]--
            dis[nxt] = max(dis[nxt], dis[top] + 1)
            if in[nxt] == 0 {
                q = append(q, nxt)
            }
        }
    }
    
    res := 0
    for i := 0; i < n; i++ {
        res += dis[i]
    }
    return res
    
}