func fullBloomFlowers(a [][]int, persons []int) []int {
    res := make([]int, len(persons));
    
    intervals := make([][]int, 0)
    pe := make([][]int, 0)
    for i := 0; i < len(a); i++ {
        var start, end int = a[i][0], a[i][1] + 1 
        intervals = append(intervals, []int{start, 1});
        intervals = append(intervals, []int{end, -1});
    }
    
    for i := 0; i < len(persons); i++ {
        pe = append(pe, []int{persons[i], i});
    }
     
    sort.SliceStable(intervals, func(i, j int) bool {
        return intervals[i][0] < intervals[j][0]
    })
    sort.SliceStable(pe, func(i, j int) bool {
        return pe[i][0] < pe[j][0]
    })
    
    
    sum := 0
    k := 0
    for i := 0; i < len(intervals); i++ {
        for k < len(pe) && pe[k][0] <= intervals[i][0] - 1 {
            res[pe[k][1]] = sum
            k++
        }
        j := i
        for j < len(intervals) && intervals[j][0] == intervals[i][0] {
            sum += intervals[j][1]
            j++
        }
        i = j - 1
    }
    
    
    return res
}