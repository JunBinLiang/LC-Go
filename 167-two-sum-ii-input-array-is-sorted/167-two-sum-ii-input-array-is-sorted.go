func twoSum(a []int, target int) []int {
    var l, r int = 0, len(a) - 1
    for l < r {
        if a[l] + a[r] == target {
            return []int{l + 1, r + 1}
        } else if a[l] + a[r] > target {
            r--
        } else {
            l++
        }
    }
    return nil
}