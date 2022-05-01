func countDistinct(a []int, k int, p int) int {
    f := make(map[string]bool)
    for i := 0; i < len(a); i++ {
        cnt := 0
        var sb strings.Builder
        for j := i; j < len(a); j++ {
            if a[j] % p == 0 {
                cnt++
            }
            if cnt > k {
                break
            }
            sb.WriteString(string(a[j]))
            sb.WriteString(",")
            f[sb.String()] = true
        }
    }
    return len(f)
}