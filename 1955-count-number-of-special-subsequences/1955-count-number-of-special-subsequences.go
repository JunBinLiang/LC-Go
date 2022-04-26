func countSpecialSubsequences(nums []int) int {
    mod := 1000000007
    var a, ab, abc int = 0, 0, 0
    for i := 0; i < len(nums); i++ {
        if nums[i] == 0 {
            a = a * 2 + 1
            a %= mod
        } else if nums[i] == 1 {
            ab = ab * 2 + a
            ab %= mod
        } else {
            abc = abc * 2 + ab
            abc %= mod
        }
    }

    return abc
}