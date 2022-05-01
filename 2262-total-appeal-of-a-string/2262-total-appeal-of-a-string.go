func atmost(s string, k int) int64 {
    cnt := make([]int, 26)
    var SIZE, j int = 0, 0
    res := int64(0)
    for i := 0; i < len(s); i++ {
        cnt[s[i] - 'a']++
        if cnt[s[i] - 'a'] == 1 {
            SIZE++
        }
        for SIZE > k {
            cnt[s[j] - 'a']--
            if cnt[s[j] - 'a'] == 0 {
                SIZE--
            }
            j++
        }
        res += int64(i - j + 1)
    }
    return res
}

func appealSum(s string) int64 {
    res := int64(0)
    for i := 1; i <= 26; i++ {
        x := atmost(s, i) - atmost(s, i - 1)
        res += (x) * int64(i)
    } 
    return res
}