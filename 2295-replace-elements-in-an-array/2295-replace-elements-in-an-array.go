func arrayChange(nums []int, operations [][]int) []int {
    f := make(map[int]int)
    for i := 0; i < len(nums); i++ {
        f[nums[i]] = i
    }
    
    for i := 0; i < len(operations); i++ {
        x := operations[i][0]
        y := operations[i][1]
        idx := f[x]
        f[x] = -1
        f[y] = idx
    }
    
    for k, v := range f {
        if v != -1 {
            nums[v] = k
        }
    }
    return nums
}