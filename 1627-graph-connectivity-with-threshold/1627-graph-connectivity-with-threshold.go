func find(nums []int, x int) int {//union find => find method
    if nums[x] == x {
       return x
    }
    root := find(nums,nums[x])
    nums[x] = root
    return root
}

func areConnected(n int, threshold int, queries [][]int) []bool {
    res := make([]bool, len(queries))
    nums := make([]int, n + 1)
    for i := 0; i < len(nums); i++ {
        nums[i] = i
    }
    
    for i := 1; i <= n; i++ {
        for j := 1; j * j <= i; j++ {
            if i % j == 0 {
                a := j
                b := i / j
                if a > threshold {
                    r1 := find(nums, i)
                    r2 := find(nums, a)
                    if r1 != r2 {
                        nums[r2] = r1
                    }
                }
                
                if b > threshold {
                    r1 := find(nums, i)
                    r2 := find(nums, b)
                    if r1 != r2 {
                        nums[r2] = r1
                    }
                }
            }
        }
    }
    
    for i := 0; i < len(queries); i++ {
        var x, y int = queries[i][0], queries[i][1]
        r1 := find(nums, x)
        r2 := find(nums, y)
        if r1 == r2 {
            res[i] = true
        } else {
            res[i] = false
        }
    }
    
    return res
}