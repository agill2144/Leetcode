// eval each block
// as-in, how much water can be blocked on top of $this block
// for this, we need 2 things
// 1. do we have a left wall ?
// 2. do we have a right wall ?
// if yes, then take the min of the 2
// and remove current block's height, thats how much we can trap
// if current block height exceeds min of #1 and #2, then we will get negative,
// meaning we cannot trap water on top of this block
// tc = o(3n) = o(n)
// sc = o(2n) = o(n)
func trap(height []int) int {
    n := len(height)
    maxL := make([]int, n)
    maxR := make([]int, n)
    for i := 1; i < n-1; i++ {maxL[i] = max(maxL[i-1], height[i-1])}
    for i := n-2; i >= 1; i-- {maxR[i] = max(maxR[i+1], height[i+1])}    
    trap := 0
    for i := 1; i < n-1; i++ {
        total := min(maxL[i], maxR[i]) - height[i]
        if total < 0 {continue}
        trap += total
    }
    return trap
}