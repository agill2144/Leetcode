func maxPoints(points [][]int) int64 {

    m := len(points)
    n := len(points[0])
    
    for i := m-2; i >= 0; i-- {
        for j := 0; j < n; j++ {
            maxPoints := 0
            for k := 0; k < n; k++ {
                maxPoints = max(maxPoints, (points[i+1][k]+points[i][j])-abs(j-k))
            }
            points[i][j] = maxPoints
        }
    }
    out := math.MinInt64
    for j := 0; j < n; j++ {
        out = max(points[0][j],out)
    }
    return int64(out)
}

func max(x, y int) int {
    if x > y {return x}
    return y
}

func abs(x int) int {
    if x < 0 {return x * -1}
    return x
}