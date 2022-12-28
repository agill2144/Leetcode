func leastBricks(wall [][]int) int {
    m := len(wall)
    maxGapCount := 0
    gapCount := map[int]int{}
    
    for i := 0; i < m; i++ {
        row := wall[i]
        gapSum := 0
        // exclude last position
        for j := 0; j < len(row)-1; j++ {
            gapSum += row[j]
            gapCount[gapSum]++
            if val := gapCount[gapSum]; val > maxGapCount {maxGapCount = val}
        }
    }
    
    return m-maxGapCount
}

