func lengthOfLongestSubstring(s string) int {
    var res, j int = 0, 0
    cnt := make([]int, 256)
    for i := 0; i < len(s); i++ {
        cnt[s[i]]++
        for cnt[s[i]] > 1 {
            cnt[s[j]]--
            j++
        }
        if i - j + 1 > res {
            res = i - j + 1
        }
    }
    return res
}