func trap(height []int) int {
    n := len(height)
    slow := 0
    fast := slow+1
    total := 0
    trap := 0
    for fast < len(height) {
        if height[slow] > height[fast] {
            trap += height[slow]-height[fast]
        } else {
            total += trap
            trap = 0
            slow = fast
        }
        fast++
    }
    peek := slow
    slow = n-1
    fast = slow-1
    trap = 0
    for fast >= peek {
        if height[slow] > height[fast] {
            trap += height[slow]-height[fast]
        } else {
            slow = fast
            total += trap
            trap = 0
        }
        fast--
    }
    return total
}