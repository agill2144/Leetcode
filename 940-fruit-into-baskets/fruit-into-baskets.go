func totalFruit(fruits []int) int {
    idx := map[int]int{}
    left := 0
    maxWin := 0
    for i := 0; i < len(fruits); i++ {
        idx[fruits[i]] = i
        if len(idx) > 2 {
            // invalid window
            // to make it valid, escape the earliest idx to keep window size large
            minIdx := len(fruits)+1
            minIdxFruit := -1
            for k, v := range idx {
                if v < minIdx {
                    minIdx = v
                    minIdxFruit = k
                }
            }
            // make sure left ptr has escaped the minIdx to make window valid
            if left <= minIdx {left = minIdx+1}
            // this fruit has left our window state
            // therefore can be removed
            delete(idx, minIdxFruit)
        }
        // now valid
        maxWin = max(maxWin, i-left+1)
    }
    return maxWin
}