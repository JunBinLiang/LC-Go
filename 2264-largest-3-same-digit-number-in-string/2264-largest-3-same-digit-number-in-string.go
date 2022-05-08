func max(a int, b int) int {
    if a > b {
        return a
    }
    return b
}

func format(num int) string {
    var res strings.Builder
    for i := 0; i < 3; i++ {
        res.WriteString(string(byte(num + '0')))
    }
    return res.String()
}

func largestGoodInteger(s string) string {
    num := -1
    for i := 0; i < len(s) - 2; i++ {
        if s[i] == s[i + 1] && s[i] == s[i + 2] {
            num = max(num, int(s[i] - '0'))
        }
    }
    if num == -1 {
        return ""
    }
    return format(num)
}