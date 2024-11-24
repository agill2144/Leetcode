func shipWithinDays(weights []int, days int) int {
    n := len(weights)
    start := math.MinInt64
    end := 0
    for i := 0; i < n; i++ {
        start = max(start, weights[i])
        end += weights[i]
    }
    ans := 0
    for start <= end {
        atMax := start + (end-start)/2
        currW := 0
        countDays := 1
        for j := 0; j < n; j++ {
            currW += weights[j]
            if currW > atMax {
                countDays++
                currW = weights[j]
            }
        }
        if countDays > days {
            start = atMax+1
        } else {
            ans = atMax
            end = atMax-1
        }
    }
    return ans
}