func twoSum(nums []int, target int) []int {
    ma := make(map[int]int)
    for i := 0; i < len(nums); i++ {
        if _, ok := ma[target - nums[i]]; ok {
            return []int{ma[target - nums[i]], i}
        }
        ma[nums[i]] = i
    }
    return nil
}