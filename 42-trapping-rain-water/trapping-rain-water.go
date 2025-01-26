func trap(height []int) int {
    n := len(height)
    maxL := make([]int, n)
    maxR := make([]int, n)
    for i := 1; i < n-1; i++ {maxL[i] = max(maxL[i-1], height[i-1])}
    for i := n-2; i >= 1; i-- {maxR[i] = max(maxR[i+1], height[i+1])}
    
    trap := 0
    for i := 1; i < n-1; i++ {
        total := min(maxL[i], maxR[i]) - height[i]
        if total < 0 {continue}
        trap += total
    }
    return trap
}