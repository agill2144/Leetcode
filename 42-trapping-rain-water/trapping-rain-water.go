func trap(height []int) int {
    n := len(height)
    leftWall := height[0]
    rightWall := height[n-1]
    left, right := 0, n-1
    total := 0
    for left < right {
        if leftWall <= rightWall {
            // we have a rightWall that is >= leftWall
            // therefore we can process the left building
            if leftWall >= height[left]{
                // we can only process left build
                // if leftWall height >= left build
                // why not just > , why >= ?
                // well 1. we have picked a side to process becuase leftWall is smaller, 
                // and rightWall provides us a guarantee that we can trap something on left, we have enough height on right to trap
                // 2. if building height == leftWall, this building cannot trap any water, even tho we have something on right side bigger to hold
                // from overflowing on right side
                // but we have nothing on left side from overflowing
                // therefore still process this left building even if height == leftWall
                // therefore >= instead of just > 
                trap := leftWall-height[left]
                total += trap
                left++
            } else {
                leftWall = height[left]
            }
        } else {
            // we have a leftWall > rightWall
            // therefore we can process the right building
            if rightWall >= height[right] {
                trap := rightWall-height[right]
                total += trap
                right--
            } else {
                rightWall = height[right]
            }
        }
    }
    return total
}

// func trap(height []int) int {
//     n := len(height)
//     maxL := make([]int, n)
//     maxR := make([]int, n)
//     for i := 1; i < n; i++ {
//         maxL[i] = max(maxL[i-1], height[i-1])
//         maxR[n-1-i] = max(maxR[n-1-i+1], height[n-1-i+1])
//     }

//     out := 0
//     for i := 1; i <= n-2; i++ {
//         left := maxL[i]
//         right := maxR[i]
//         trap := min(left, right) - height[i]
//         if trap > 0 {out += trap}
//     }
//     return out
// }