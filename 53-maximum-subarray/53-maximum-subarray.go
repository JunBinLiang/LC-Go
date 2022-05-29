func min(a int, b int) int {
    if a < b {
        return a
    }    
    return b
}

func max(a int, b int) int {
    if a > b {
        return a
    }
    return b
}

func maxSubArray(nums []int) int {
    var res, sum, mn = -1000000000, 0, 0
    for i := 0; i < len(nums); i++ {
        sum += nums[i]
        res = max(res, sum - mn);
        mn = min(mn, sum)
    }
    return res
}