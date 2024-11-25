/*
    approach: inverted sliding window
    - giving edgy sliding window vibes
    - where a window wraps around as if the array was circular
        - when you see this, try to think of inverted window
        - inverted window items outside of a sliding window
        - where the current window is not necessarily the answer
        - but the elements outside of the window could be
        - inverted sliding windows are specially helpful when a window needs to wrap around the array
        - [x,x,x,x,x,x,x,x]
               |_______|
        - if the above is our current window, the elements outside are automatically wrapping!
            - thereby creating the wrapping window outside of current regular sliding window
        - [x,x,x,x,x,x,x,x]
           |_____|
        - [x,x,x,x,x,x,x,x]
               |_____|
        - [x,x,x,x,x,x,x,x]
                   |_____|
    - we need to reduce x to 0
    - and these elements can only be picked from the outside edges of the array
        - i.e elements that are part of the inverted window
    - we can create a total sum of nums array
    - and can remove x from total sum
    - this means, if we creating a window with the total-x sum, we will have exactly x sum in the inverted window
    time = o(n)
    space = o(1)
*/
func minOperations(nums []int, x int) int {
    n := len(nums)
    k := 0
    for i := 0; i < n; i++ {k += nums[i]}
    k -= x
    left := 0
    rSum := 0
    minWin := len(nums)+1
    for i := 0; i < n; i++ {
        rSum += nums[i]
        for rSum > k && left <= i {
            rSum -= nums[left]
            left++
        }
        if rSum == k {
            minWin = min(minWin, n-(i-left+1))            
        }
    }
    if minWin == len(nums)+1 {minWin = -1}
    return minWin 
}