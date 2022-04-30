func isPrefix(s string, t string) bool {
    if len(s) > len(t) {
        return false
    }
    for i := 0; i < len(s); i++ {
        if(s[i] != t[i]) {
            return false
        }
    }
    return true
}

func countPrefixes(a []string, s string) int {
    cnt := 0
    for i := 0; i < len(a); i++ {
        if isPrefix(a[i], s) {
            cnt++
        }
    }
    return cnt
}