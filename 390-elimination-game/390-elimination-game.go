func lastRemaining(n int) int {
    a := make([]int, 0)
    cnt := 0
    res := 1
    for n > 1 {
        if n % 2 == 1 {
            a = append(a, 1)
            res = 2
        } else {
            if cnt % 2 == 0 {
                a = append(a, 1)
                res = 2
            } else {
                a = append(a, 0)
                res = 1
            }
        }
        cnt++
        n /= 2
    }
    
    //fmt.Println(a)
    
    for i := len(a) - 2; i >= 0; i-- {
        if a[i] == 1 {
            res *= 2
        } else {
            res = res *  2 - 1
        }
    }
    return res
}