func finalValueAfterOperations(a []string) int {
    res := 0
    for i := 0; i < len(a); i++ {
        s := a[i]
        if s[0] == 'X' {
            if s[1] == '+' {
                res++
            } else {
                res--
            }
        } else {
            if s[0] == '+' {
                res++
            } else {
                res--
            }
        }
    }
    return res
}