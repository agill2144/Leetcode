func leastBricks(wall [][]int) int {
    gapCount := map[int]int{}
    m := len(wall)
    highestGapCount := 0
    for i := 0; i < m; i++ {
        gapTotal := 0
        for j := 0; j < len(wall[i])-1; j++ {
            gapTotal += wall[i][j]
            gapCount[gapTotal]++
            if count := gapCount[gapTotal]; count > highestGapCount {
                highestGapCount = count
            }
        }
    }
    return m-highestGapCount
}