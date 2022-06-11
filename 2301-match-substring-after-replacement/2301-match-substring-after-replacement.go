func matchReplacement(s string, sub string, mappings [][]byte) bool {
    ok := make([][]bool, 256)
    for i := 0; i < 256; i++ {
        ok[i] = make([]bool, 256)
    }
    
    for i := 0; i < len(mappings); i++ {
        c1 := mappings[i][0]
        c2 := mappings[i][1]
        ok[int(c1)][int(c2)] = true
    }
    
    for i := 0; i < len(s); i++ {
        if i + len(sub) - 1 >= len(s) {
            break
        }
        
        good := true
        for j := 0; j < len(sub); j++ {
            if sub[j] != s[i + j] {
                c1 := int(sub[j])
                c2 := int(s[i + j])
                if !ok[c1][c2] {
                    good = false
                    break
                }
            } 
        }
        if good {
            return true
        }
    }
    
    return false
}