func find(nums []int, x int) int {//union find => find method
    if nums[x] == x {
       return x
    }
    root := find(nums,nums[x])
    nums[x] = root
    return root
}

func max(a int, b int) int {
    if a > b {
        return a
    }
    return b
}

func gcdSort(a []int) bool {
    mx := 0
    for i := 0; i < len(a); i++ {
        mx = max(mx, a[i])
    }
    
    is := make([]bool, mx + 1)
    for i := 0; i < len(a); i++ {
        is[a[i]] = true
    }
    
    nums := make([]int, mx + 1)
    for i := 0; i < len(nums); i++ {
        nums[i] = i
    }
    
    for i := 2; i <= mx; i++ {
        for j := i; j <= mx; j += i {
            if is[j] {
                r1 := find(nums, i)
                r2 := find(nums, j)
                if r1 != r2 {
                    nums[r1] = r2
                }
            }
        } 
    }
    
    f := make(map[int][]int)
    for i := 0; i < len(a); i++ {
        r := find(nums, a[i])
        f[r] = append(f[r], a[i])
    }
    
    for _, v := range f {
        sort.Ints(v)
    }
    
    
    b := make([]int, len(a))
    for i := 0; i < len(a); i++ {
        r := find(nums, a[i])
        b[i] = f[r][0]
        f[r] = f[r][ 1 : ]
    }
    
    
    for i := 1; i < len(b); i++ {
        if b[i] < b[i - 1] {
            return false
        }
    }
    return true
    
}