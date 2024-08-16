func maxDistance(arrays [][]int) int {
    n := len(arrays)
    minVal := arrays[0][0]
    maxVal := arrays[0][len(arrays[0])-1]
    res := math.MinInt64
    for i := 1; i < n; i++ {
        currSmall := arrays[i][0]
        currLarge := arrays[i][len(arrays[i])-1]
        res = max(res, max(currLarge-minVal, maxVal-currSmall))
        minVal = min(minVal, currSmall)
        maxVal = max(maxVal, currLarge)
    }
    return res
}

