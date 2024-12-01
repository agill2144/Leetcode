func trap(height []int) int {
    n := len(height)
    left := make([]int, n)
    right := make([]int, n)
    for i := 1; i < n-1; i++ {left[i] = max(left[i-1], height[i-1])}
    for i := n-2; i >= 0; i-- { right[i] = max(right[i+1], height[i+1]) }

    total := 0
    for i := 1; i < n-1; i++ {
        curr := height[i]
        unit := min(left[i], right[i]) - curr
        if unit > 0 {
            total += unit
        }
    }

    return total
}