func trap(height []int) int {
    n := len(height)
    maxL := make([]int, n)
    maxR := make([]int, n)
    for i := 0; i < n; i++ {
        if i != 0 {
            maxL[i] = max(maxL[i-1], height[i-1])
        }
        idx := n-1-i
        if idx != n-1 {
            maxR[idx] = max(maxR[idx+1], height[idx+1])
        }
    }
    total := 0
    for i := 0; i < n; i++ {
        curr := height[i]
        trap := min(maxL[i], maxR[i])-curr
        if trap > 0 {
            total += trap
        }
    }
    return total
}