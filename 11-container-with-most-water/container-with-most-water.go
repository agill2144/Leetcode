func maxArea(height []int) int {
    left := 0
    right := len(height)-1
    maxRes := 0
    for left < right {
        leftH := height[left]
        rightH := height[right]
        width := right-left
        area := min(leftH, rightH) * width
        maxRes = max(maxRes, area)
        if leftH < rightH {left++} else {right--}
    }
    return maxRes
}