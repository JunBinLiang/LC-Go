func min(a int, b int) int {
    if a < b {
        return a
    }
    return b
}

func maxCount(m int, n int, ops [][]int) int {
    var a, b int = m, n
    for i := 0; i < len(ops); i++ {
        a = min(a, ops[i][0])
        b = min(b, ops[i][1])
    }
    return a * b
}