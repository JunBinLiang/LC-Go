func removeDuplicates(s string, k int) string {
    sta := make([][]int, 0)
    for i := 0; i < len(s); i++ {
        if len(sta) == 0 || int(s[i]) != sta[len(sta) - 1][0] {
            sta = append(sta, []int{int(s[i]), 1})
        } else {
            sta[len(sta) - 1][1]++
            if sta[len(sta) - 1][1] == k {
                sta = sta[ : len(sta) - 1]
            }
        }
    }
    
    var res strings.Builder
    for i := 0; i < len(sta); i++ {
        var c, t int = sta[i][0], sta[i][1]
        for t > 0 {
            t--
            res.WriteString(string(byte(c)))
        }
    }
    return res.String()
}