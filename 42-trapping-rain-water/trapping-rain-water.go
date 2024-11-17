func trap(height []int) int {
    n := len(height)
    maxL := make([]int, n)    
    maxR := make([]int, n)
    for i := 1; i < n; i++ {maxL[i] = max(height[i-1], maxL[i-1])}
    for i := n-2; i >= 0; i-- {maxR[i] = max(height[i+1], maxR[i+1])}

    total := 0
    for i := 1; i < n-1; i++ {
        h := height[i]
        // we can trap on top of this
        // if there is a height on left and right greater than itself
        lh := maxL[i]
        rh := maxR[i]
        if lh > h && rh > h {
            total += min(lh,rh)-h
        }
    }
    return total
    
}