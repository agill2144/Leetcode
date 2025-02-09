func trap(height []int) int {
    n := len(height)
    leftWall := 0
    left := 1
    rightWall := n-1
    right := n-2
    total := 0
    for left <= right {
        if height[rightWall] >= height[leftWall] {
            if height[left] <= height[leftWall]{
                total += (height[leftWall] - height[left])
                left++
            } else {
                leftWall = left
            }
        } else {
            if height[right] <= height[rightWall] {
                total += (height[rightWall] - height[right])
                right--
            } else {
                rightWall = right
            }
        }
    }
    return total
}