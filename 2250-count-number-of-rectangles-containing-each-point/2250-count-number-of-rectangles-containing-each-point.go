func countRectangles(a [][]int, points [][]int) []int {
    res := make([]int, len(points))
    
    f := make(map[int][]int)
    for i := 0; i < len(a); i++ {
        f[a[i][1]] = append(f[a[i][1]], a[i][0])
    }
    
    for _, v := range f {
        sort.Ints(v)
    }
    
    
    for i := 0; i < len(points); i++ {
        var x, y int = points[i][0], points[i][1] 
        cnt := 0
        for h := y; h <= 100; h++ {
            if _, ok := f[h]; ok {
                var l, r, pos int = 0, len(f[h]) - 1, -1
                for l <= r {
                    mid := l + (r - l) / 2
                    if f[h][mid] >= x {
                        pos = mid
                        r = mid - 1
                    } else {
                        l = mid + 1
                    }
                }
                if pos != -1 {
                    size := len(f[h])
                    cnt += (size - pos)
                }
            }
        }
        res[i] = cnt
    }
    
    return res
}