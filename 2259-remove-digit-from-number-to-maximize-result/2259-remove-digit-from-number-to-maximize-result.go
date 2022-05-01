func removeDigit(s string, digit byte) string {
    res := make([]string, 0)
    for i := 0; i < len(s); i++ {
        if s[i] == digit {
            var sb strings.Builder
            for j := 0; j < len(s); j++ {
                if j == i {
                    continue
                }
                sb.WriteString(string(s[j]))
            }
            res = append(res, sb.String())
        }
    }
    sort.Strings(res)
    return res[len(res) - 1]
}