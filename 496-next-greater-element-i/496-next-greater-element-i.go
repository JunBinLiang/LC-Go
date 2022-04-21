func nextGreaterElement(nums1 []int, nums2 []int) []int {
    f := make(map[int]int)
    res := make([]int, len(nums1))
    for i := 0; i < len(nums1); i++ {
        f[nums1[i]] = i
        res[i] = -1
    }
    
    sta := make([]int, 0)
    
    for i := 0; i < len(nums2); i++ {
        for len(sta) > 0 && nums2[i] > nums1[sta[len(sta) - 1]] {
            res[sta[len(sta) - 1]] = nums2[i]
            sta = sta[ : len(sta) - 1]
        }
        
        if _, ok := f[nums2[i]]; ok {
            sta = append(sta, f[nums2[i]])
        }
    }
    
    return res
}