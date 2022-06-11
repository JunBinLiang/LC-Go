func strongPasswordCheckerII(s string) bool {
    if len(s) < 8 {
        return false
    }
    
    for i := 1; i < len(s); i++ {
        if s[i] == s[i - 1] {
            return false
        }
    }
    
    spe := "!@#$%^&*()-+"
    
    var a1, a2, a3, a4 bool = false, false, false, false
    for i := 0; i < len(s); i++ {
        for j := 'a'; j <= 'z'; j++ {
            if s[i] == byte(j) {
                a1 = true
            }
        }
        
        for j := 'A'; j <= 'Z'; j++ {
            if s[i] == byte(j) {
                a2 = true
            }
        }
        
        for j := 0; j < len(spe); j++ {
            if spe[j] == s[i] {
                a3 = true
            }
        }
        
        for j := '0'; j <= '9'; j++ {
            if s[i] == byte(j) {
                a4 = true
            }
        }
        
    }
    
    return a1 && a2 && a3 && a4
}