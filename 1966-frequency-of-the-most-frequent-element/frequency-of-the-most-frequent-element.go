/*
    greedy / sliding window from the back
    - we can only increment a number at a idx!
    - [1,2,3,4]
    - if we increment highest val (4), we cannot match with any other number..
    - therefore we will never increment the highest number
    - we will infact, try to match smaller adj value to the highest number
    - therefore lets sort first, so we have the highest in the back of the array
    - now we will loop from n-2
    - now we have to try and match n-2 (3) with n-1 (4)
    - if we could match the sum of these 2 value would be 8
        - n-2 becomes 4 and n-1 is already 4
    - if they matched, their window sum is supposed to be 8
        - window size * n-1
    - how far is our current window sum
        - (n-1) + (n-2) = 3+4 = 7
    - we are off by 1, and now if our k <= 1; 
    - this is a valid window!
    - if its a valid window, expand; move i back; keep right ptr on n-1
    - if its a invalid window, no point in expanding;
        - if desired sum of windowSize 2 could not be met, adding another element to this window will
        - only increase the window sum, therefore no need to expand, 
        - instead SHRINK from right
        - right--; update our runningSum by excluding the right value leaving our window
*/
func maxFrequency(nums []int, k int) int {
    sort.Ints(nums)
    n := len(nums)
    maxWin := 1
    right := n-1
    rSum := nums[n-1]
    for i := n-2; i >= 0; i-- {
        size := right-i+1
        val := nums[right]
        desiredSum := val * size
        rSum += nums[i]
        if desiredSum - rSum <= k {
            maxWin = max(maxWin, size)
            continue
        } else {
            rSum -= nums[right]
            right--
        }
    }
    return maxWin
}