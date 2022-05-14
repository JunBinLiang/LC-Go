func waysToSplitArray(nums []int) int {
    var sum, l int64 = 0, 0
    res := 0
    for i := 0; i < len(nums); i++ {
        sum += int64(nums[i])
    }
    for i := 0; i < len(nums) - 1; i++ {
        l += int64(nums[i])
        sum -= int64(nums[i])
        if l >= sum {
            res++
        }
    }
    return res
}