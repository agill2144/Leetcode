func trap(height []int) int {
    n := len(height)
    lh := 0
    rh := n-1
    l := 0
    r := n-1
    total := 0
    for l < r {
        if height[rh] >= height[lh] {
            // process left side
            if height[l] <= height[lh] {
                total += (height[lh]-height[l])
                l++
            } else {
                lh = l
            }
        } else {
            if height[r] <= height[rh] {
                total += (height[rh]-height[r])
                r--
            } else {
                rh = r
            }
        }
    }
    return total
}
// func trap(height []int) int {
//     n := len(height)
//     maxL := make([]int, n)
//     maxR := make([]int, n)
//     for i := 0; i < n; i++ {
//         if i != 0 {
//             maxL[i] = max(maxL[i-1], height[i-1])
//         }
//         idx := n-1-i
//         if idx != n-1 {
//             maxR[idx] = max(maxR[idx+1], height[idx+1])
//         }
//     }
//     total := 0
//     for i := 0; i < n; i++ {
//         curr := height[i]
//         trap := min(maxL[i], maxR[i])-curr
//         if trap > 0 {
//             total += trap
//         }
//     }
//     return total
// }