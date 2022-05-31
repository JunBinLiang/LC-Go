func hasAllCodes(s string, k int) bool {
    set := make(map[int]int)
    for i := 0; i < len(s); i++ {
        if i + k - 1 >= len(s) {
            break
        }
        val := 0
        for j := i; j < i + k; j++ {
            val = val * 2
            val += (1 << int(s[j] - '0'))
        }
        set[val] = 1
    }
    return len(set) == (1 << k);
}