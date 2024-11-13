func trap(height []int) int {
    n := len(height)
    maxL := make([]int, n)
    maxR := make([]int, n)
    for i := 1; i < n; i++ {maxL[i] = max(maxL[i-1], height[i-1])}
    for i := n-2; i >= 0; i-- {maxR[i] = max(maxR[i+1], height[i+1])}
    total := 0
    for i := 1; i < n-1; i++ {
        units := min(maxL[i], maxR[i]) - height[i]
        if units > 0 {total += units}
    }
    return total
}

