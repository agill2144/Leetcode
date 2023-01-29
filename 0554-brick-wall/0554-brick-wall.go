func leastBricks(wall [][]int) int {
    gapCount := map[int]int{}
    maxCount := 0
    m := len(wall)
    for i := 0; i < m; i++ {
        n := len(wall[i])
        sum := 0
        for j := 0; j < n-1; j++ {
            sum += wall[i][j]
            gapCount[sum]++
            if gapCount[sum] > maxCount {maxCount = gapCount[sum]}
        }
    }
    return m-maxCount
}