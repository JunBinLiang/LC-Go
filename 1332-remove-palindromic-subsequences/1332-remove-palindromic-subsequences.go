func isp(s string) bool {
    var l, r int = 0, len(s) - 1
    for l < r {
        if s[l] != s[r] {
            return false
        }
        l++
        r--
    }
    return true
}

func removePalindromeSub(s string) int {
    if isp(s) {
        return 1
    }
    return 2
}