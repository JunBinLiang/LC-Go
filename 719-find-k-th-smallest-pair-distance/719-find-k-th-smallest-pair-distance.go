func get(a []int, mid int) int {
    cnt := 0
    j := 0
    for i := 0; i < len(a); i++ {
        for a[i] - a[j] > mid {
            j++
        }
        cnt += (i - j)
    }
    return cnt
}

func smallestDistancePair(a []int, k int) int {
    n := len(a)
    sort.Ints(a)
    var l, r, res int = 0, a[n - 1] - a[0], -1
    for l <= r {
        mid := l + (r - l) / 2
        cnt := get(a, mid)
        if cnt >= k {
            res = mid
            r = mid - 1
        } else {
            l = mid + 1
        }
    }
    return res
}