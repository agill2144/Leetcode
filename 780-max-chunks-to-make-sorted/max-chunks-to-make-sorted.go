func maxChunksToSorted(arr []int) int {
    count := 0
    maxVal := math.MinInt64
    for i := 0; i < len(arr); i++ {
        maxVal = max(maxVal, arr[i])
        if i == maxVal {count++; maxVal = math.MinInt64}
    }
    return count
}