func numberOfSteps(num int) int {
    res := 0
    for num > 0 {
        res++
        if num % 2 == 0 {
            num /= 2
        } else {
            num--
        }
    }
    return res
}