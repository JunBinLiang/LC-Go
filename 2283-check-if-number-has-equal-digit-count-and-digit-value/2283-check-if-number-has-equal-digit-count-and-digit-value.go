func digitCount(s string) bool {
    cnt := make([]int, 10)
    for i := 0; i < len(s); i++ {
        cnt[int(s[i] - '0')]++
    }
    
    for i := 0; i < len(s); i++ {
        if cnt[i] != int(s[i] - '0') {
            return false
        }
    }
    
    return true
}