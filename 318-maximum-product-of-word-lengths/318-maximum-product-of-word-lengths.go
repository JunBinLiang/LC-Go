func max(a int, b int) int {
    if a > b {
        return a
    }
    return b
}

func maxProduct(words []string) int {
    bits := make([]int, len(words))
    for i := 0; i < len(bits); i++ {
        for j := 0; j < len(words[i]); j++ {
            bits[i] |= (1 << (int(words[i][j] - 'a')))
        }
    }
    
    res := 0
    for i := 0; i < len(words); i++ {
        for j := i + 1; j < len(words); j++ {
            if (bits[i] & bits[j]) == 0 {
                res = max(res, len(words[i]) * len(words[j]))
            }
        }
    }
    
    return res
}