func find(nums []int, x int) int {//union find => find method
    if nums[x] == x {
       return x
    }
    root := find(nums,nums[x])
    nums[x] = root
    return root
}

func smallestStringWithSwaps(s string, pairs [][]int) string {
    var sb strings.Builder
    nums := make([]int, len(s))
    for i := 0; i < len(nums); i++ {
        nums[i] = i
    }
    
    for i := 0; i < len(pairs); i++ {
        var u, v int = pairs[i][0], pairs[i][1]
        r1 := find(nums, u)
        r2 := find(nums, v)
        if r1 != r2 {
            nums[r1] = r2
        }
    }
    
    f := make(map[int][]int)
    for i := 0; i < len(s); i++ {
        r := find(nums, i)
        f[r] = append(f[r], int(s[i]))
    }
    
    for _, v := range f {
        sort.Ints(v)
    }
    
    for i := 0; i < len(s); i++ {
        r := find(nums, i)
        sb.WriteString(string(byte(f[r][0])))
        f[r] = f[r][1 : ]
    }
    
    return sb.String()
}