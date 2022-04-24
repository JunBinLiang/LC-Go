func inside(x1 int, y1 int, x2 int, y2 int, r int) bool {
    return r * r >= ((x1 - x2)) * (x1 - x2) + (y1 - y2) * (y1 - y2)
}

func countLatticePoints(circles [][]int) int {
    cnt := 0
    for x := 0; x <= 400 ; x++ {
        for y := 0; y <= 400; y++ {
            for i := 0; i < len(circles); i++ {
            if inside(x, y, circles[i][0], circles[i][1], circles[i][2]){
                    cnt++
                    break
                }
            }
        } 
    }
    
    return cnt
}