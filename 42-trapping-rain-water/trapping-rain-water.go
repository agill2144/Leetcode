func trap(height []int) int {
    n := len(height)
    leftWall := 0
    rightWall := n-1
    left, right := 0, n-1
    total := 0
    for left < right {
        if height[rightWall] >= height[leftWall] {
            if height[leftWall] >= height[left] {
                total += (height[leftWall]-height[left])
                left++
            } else {
                leftWall = left
            }
        } else {
            if height[rightWall] >= height[right] {
                total += (height[rightWall]-height[right])
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
//     // space = o(n)
//     maxL := make([]int, n)
//     // space = o(n)
//     maxR := make([]int, n)
//     // time = o(2n)
//     for i := 1; i < n; i++ {maxL[i] = max(height[i-1], maxL[i-1])}
//     for i := n-2; i >= 0; i-- {maxR[i] = max(height[i+1], maxR[i+1])}

//     total := 0
//     // time = o(n)
//     // total = o(3n) or o(n)
//     // space = o(2n) or o(n)
//     for i := 1; i < n-1; i++ {
//         h := height[i]
//         // we can trap on top of this
//         // if there is a height on left and right greater than itself
//         lh := maxL[i]
//         rh := maxR[i]
//         if lh > h && rh > h {
//             total += min(lh,rh)-h
//         }
//     }
//     return total
// }