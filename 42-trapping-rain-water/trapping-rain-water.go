func trap(height []int) int {
    n := len(height)
    maxLeft := make([]int, n)
    maxRight := make([]int, n)
    for i := 0; i < n; i++ {
        if i == 0 || i == n-1 {continue}
        maxLeft[i] = max(maxLeft[i-1], height[i-1])
        maxRight[n-1-i] = max(maxRight[n-1-i+1], height[n-1-i+1])
    }
    total := 0
    for i := 0; i < len(height); i++ {
        ml := maxLeft[i]
        mr := maxRight[i]
        smallest := min(ml,mr)
        trapped := smallest-height[i]
        if trapped > 0 {
            total += trapped
        }
    }
    return total
}