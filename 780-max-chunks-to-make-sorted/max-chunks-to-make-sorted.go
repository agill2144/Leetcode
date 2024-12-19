func maxChunksToSorted(arr []int) int {
    maxVal := -1
    count := 0
    for i := 0; i < len(arr); i++ {
        maxVal = max(maxVal, arr[i])
        if maxVal == i {count++}
    }
    return count
}