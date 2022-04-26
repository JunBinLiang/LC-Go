func find(nums []int, x int) int {//union find => find method
    if nums[x] == x {
       return x
    }
    root := find(nums,nums[x])
    nums[x] = root
    return root
}

func hitBricks(grid [][]int, hits [][]int) []int {
    var n, m int = len(grid), len(grid[0])
    res := make([]int, len(hits))
    nums := make([]int, n * m + 1)
    cnt := make([]int, n * m + 1)
    dir := [4][2]int{{0, -1}, {-1, 0}, {0, 1}, {1, 0}}
    bad := make([]bool, len(hits))
    seen := make([]bool, n * m + 1)
    
    for i := 0; i < len(nums); i++ {
        nums[i] = i
        cnt[i] = 1
    }
    
    for i := 0; i < len(hits); i++ {
        if grid[hits[i][0]][hits[i][1]] == 0 {
            bad[i] = true
        }
        grid[hits[i][0]][hits[i][1]] = 0
    }
    
    for i := 0; i < n; i++ {
        for j := 0; j < m; j++ {
            if grid[i][j] == 0 {
                continue
            }
            id1 := i * m + j
            for x := 0; x < 2; x++ {
                var row, col int = i + dir[x][0], j + dir[x][1]
                if row < 0 || row >= n || col < 0 || col >= m || grid[row][col] == 0 {
                    continue
                }
                id2 := row * m + col
                r1 := find(nums, id1)
                r2 := find(nums, id2)
                if r1 != r2 {
                    nums[r1] = r2
                    cnt[r2] += cnt[r1]
                }
            }
        }
    }
    
    for i := 0; i < m; i++ {
        if grid[0][i] == 1 {
            r1 := find(nums, n * m)
            r2 := find(nums, i)
            if r1 != r2 {
                nums[r2] = r1
                cnt[r1] += cnt[r2]
            }
        }
    }
    
    
    for i := len(hits) - 1; i >= 0; i-- {
        if bad[i] {
            continue
        }
        
        var r, c int = hits[i][0], hits[i][1]
        grid[r][c] = 1
        connect := false
        id1 := r * m + c
        if r == 0 {
            connect = true
        }
        
        
        for x := 0; x < 4; x++ {
            var row, col int = r + dir[x][0], c + dir[x][1]
            if row < 0 || row >= n || col < 0 || col >= m || grid[row][col] == 0 {
                continue
            }
            id2 := row * m + col
            r2 := find(nums, id2)
            if r2 == n * m {
                connect = true
            }
        }
        
        
        for x := 0; x < 4; x++ {
            var row, col int = r + dir[x][0], c + dir[x][1]
            if row < 0 || row >= n || col < 0 || col >= m || grid[row][col] == 0 {
                continue
            }
            id2 := row * m + col
            r2 := find(nums, id2)
            if connect && r2 != n * m && !seen[r2] {
                res[i] += cnt[r2]
                seen[r2] = true
            }
            
        }
        
        for x := 0; x < 4; x++ {
            var row, col int = r + dir[x][0], c + dir[x][1]
            if row < 0 || row >= n || col < 0 || col >= m || grid[row][col] == 0 {
                continue
            }
            id2 := row * m + col
            r1 := find(nums, id1)
            r2 := find(nums, id2)
            
            
            seen[r2] = false
            if connect {
                if r2 != n * m {
                    nums[r2] = n * m
                    cnt[n * m] += cnt[r2]
                }
            } else {
                if r1 != r2 {
                    nums[r1] = r2
                    cnt[r2] += cnt[r1]
                }
            }
        }
        
        
        
        if connect {
            r1 := find(nums, id1)
            if r1 != n * m {
                nums[r1] = n * m
                cnt[n * m] += cnt[r1]
            }
        }
    }
    return res
}
