func jump(nums []int) int {
    if len(nums) <= 1 {return 0}
    jumps := 0
    right := 0
    i := 0
    maxRight := 0
    for i < len(nums) {
        // while processing current window [i:right]
        // find the new farthest right we can have for next window 
        // from our current window
        maxRight = max(maxRight,i + nums[i])
        // once we are at the end of current window
        // we can make our new window by updating 
        // right = maxRight that we have found so far from current window
        // and left ptr ( i ) can keep moving forward by 1
        // and because we are jumping to new window of elements
        // we increment jumps++
        if i == right {jumps++; right = maxRight}

        // if our current window is at n-1 or beyond n-1, this is it
        // our current jump count so far is our answer
        // hence break and return the jumps counted
        if right >= len(nums)-1 {
            break
        }
        i++
    }
    return jumps
}