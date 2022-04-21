func xorGame(a []int) bool {
    var xor, n int = 0, len(a)
    for i := 0; i < n; i++ {
        xor ^= a[i]
    }
    
    if xor == 0 {
        return true
    }
    
    sort.Ints(a)
    if a[0] == a[n - 1] {
        return false
    }
    
    if n % 2 == 1 {
        return false
    } 
    return true

}