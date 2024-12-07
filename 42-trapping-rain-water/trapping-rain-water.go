func trap(height []int) int {
    n := len(height)
    leftWall := 0
    rightWall := n-1
    left := 0
    right := n-1
    total := 0
    for left <= right {
        if height[rightWall] >= height[leftWall] {
            if height[leftWall] >= height[left] {
                total += height[leftWall]-height[left]
                left++
            } else {
                leftWall = left
            }
        } else {
            if height[rightWall] >= height[right] {
                total += height[rightWall] - height[right]
                right--
            } else {
                rightWall =right
            }
        }
    }
    return total
}
// func trap(height []int) int {
//     n := len(height)
//     maxL := make([]int, n)
//     maxR := make([]int, n)
//     for i := 1; i < n; i++ {maxL[i] = max(maxL[i-1], height[i-1])}
//     for i := n-2; i >=0 ; i-- {maxR[i] = max(maxR[i+1], height[i+1])}
//     total := 0
//     for i := 1; i < n-1; i++ {
//         left := maxL[i]
//         right := maxR[i]
//         trap := min(left, right)-height[i]
//         if trap > 0 {total += trap}
//     }
//     return total
// }