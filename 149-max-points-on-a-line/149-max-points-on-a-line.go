func gcd(a int, b int) int {
    if b != 0 {
        return gcd(b, a % b)
    }
    return a
}

func max(a int, b int) int {
    if a > b {
        return a
    }
    return b
}

func abs(a int) int {
    if a < 0 {
        return -a
    }
    return a
}

func maxPoints(a [][]int) int {
    res := 0
    n := len(a)
    for i := 0; i < n; i++ {
        for j := i + 1; j < n; j++ {
            x := a[i][0] - a[j][0]
            y := a[i][1] - a[j][1]
            g := gcd(abs(x), abs(y))
            
            if g == 0 {
                continue
            }
            
            cnt := 2
            for k := 0; k < n; k++ {
                if k == i || k == j {
                    continue
                }
                
                x1 := a[i][0] - a[k][0]
                y1 := a[i][1] - a[k][1]
                g1 := gcd(abs(x1), abs(y1))
                 
                if g1 == 0 {
                    continue
                }
     
                if x1 / g1 == x / g && y1 / g1 == y / g {
                    cnt++
                }
            }
            res = max(res, cnt)
        }
    }
    
    
    for i := 0; i < n; i++ {
        cnt := 1
        for j := 0; j < n; j++ {
            if i != j {
                if a[i][0] == a[j][0] {
                    cnt++
                }
            }
        }
        res = max(res, cnt)
    }
    
    for i := 0; i < n; i++ {
        cnt := 1
        for j := 0; j < n; j++ {
            if i != j {
                if a[i][1] == a[j][1] {
                    cnt++
                }
            }
        }
        res = max(res, cnt)
    }
    
    return res
}