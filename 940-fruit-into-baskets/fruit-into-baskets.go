func totalFruit(fruits []int) int {
    idx := map[int]int{}
    n := len(fruits)
    maxWindow := 0
    left := 0
    for i := 0; i < n; i++ {
        idx[fruits[i]] = i
        if len(idx) > 2 {
            // escape the earliest idx
            leftMostIdx := n+1
            for _, v := range idx {
                leftMostIdx = min(leftMostIdx, v)
            }
            if left <= leftMostIdx {
                left = leftMostIdx+1
            }
            delete(idx, fruits[leftMostIdx])
        }
        maxWindow = max(maxWindow, i-left+1)
    }
    return maxWindow
}