func is(s string, t string) bool {
    cnt := make([]int, 26)
    for i := 0; i < len(s); i++ {
        cnt[s[i] - 'a'] ++
    }
    for i := 0; i < len(t); i++ {
        cnt[t[i] - 'a'] --
    }
    
    for i := 0; i < 26; i++ {
        if cnt[i] != 0 {
            return false
        }
    }
    return true
}

func removeAnagrams(a []string) []string {
    
    sta := make([]string, 0)
    for i := 0; i < len(a); i++ {
        if len(sta) == 0 {
            sta = append(sta, a[i])
        } else {
            if is(sta[len(sta) - 1], a[i]) {
            } else {
                sta = append(sta, a[i])
            }
        }
    }
    return sta
}