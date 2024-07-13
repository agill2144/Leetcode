func trap(height []int) int {
    n := len(height)
    maxL := make([]int, n)
    maxR := make([]int, n)
    for i := 1; i < n; i++ {
        maxL[i] = max(maxL[i-1], height[i-1])
        maxR[n-1-i] = max(maxR[n-1-i+1], height[n-1-i+1])
    }

    out := 0
    for i := 1; i <= n-2; i++ {
        left := maxL[i]
        right := maxR[i]
        trap := min(left, right) - height[i]
        if trap > 0 {out += trap}
    }
    return out
}