/*

*/
func trap(height []int) int {
    n := len(height)
    left := make([]int, n)
    right := make([]int, n)
    for i := 0; i < len(height); i++ {
        if i == 0 || i == n-1 {continue}
        maxLeftSoFar := max(left[i-1], height[i-1])
        maxRightSoFar := max(right[n-1-i+1], height[n-1-i+1])
        left[i] = maxLeftSoFar
        right[n-1-i] = maxRightSoFar
    }
    total := 0
    for i := 0; i < n; i++ {
        leftH, rightH := left[i],right[i]
        canBeTrappedAboveIthBuilding := min(leftH, rightH)-height[i]
        if canBeTrappedAboveIthBuilding > 0 {
            total += canBeTrappedAboveIthBuilding
        }
    }
    return total
}