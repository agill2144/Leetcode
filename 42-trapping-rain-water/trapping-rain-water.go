func trap(height []int) int {
    n := len(height)
    leftWall := 0
    rightWall := n-1
    left := 0
    right := n-1
    total := 0
    for left < right {
        if height[rightWall] >= height[leftWall] {
            // process left side, we have a bigger right wall
            if height[leftWall] >= height[left] {
                total += max(height[leftWall], height[left]) - height[left]
                left++
            } else {
                leftWall = left
            }
        } else {
            // process right side, we have a bigger left wall
            if height[rightWall] >= height[right] {
                total += max(height[rightWall], height[right]) - height[right]
                right--
            } else {
                rightWall = right
            }
        }
    }
    return total 
}

// func trap(height []int) int {
//     n := len(height)
//     left := make([]int, n)
//     right := make([]int, n)
//     for i := 1; i < n-1; i++ {left[i] = max(left[i-1], height[i-1])}
//     for i := n-2; i >= 0; i-- { right[i] = max(right[i+1], height[i+1]) }

//     total := 0
//     for i := 1; i < n-1; i++ {
//         curr := height[i]
//         unit := min(left[i], right[i]) - curr
//         if unit > 0 {
//             total += unit
//         }
//     }

//     return total
// }