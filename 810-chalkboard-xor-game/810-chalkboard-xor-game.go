func xorGame(a []int) bool {
    var xor, n int = 0, len(a)
    for i := 0; i < n; i++ {
        xor ^= a[i]
    }
    
    if xor == 0 {
        return true
    }
    
    
    
    if n % 2 == 1 {
        return false
    } 
    return true

}