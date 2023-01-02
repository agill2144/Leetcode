func rob(nums []int) int {
    if len(nums) == 1 {return nums[0]}
    var helper func(start, end int) int
    helper = func(start, end int) int {
        prev := 0
        curr := nums[start]
        for i := start+1; i < end; i++ {
            newCurr := max(curr, nums[i]+prev)
            prev = curr
            curr = newCurr
        }
        return curr
    }
    return max(helper(0, len(nums)-1), helper(1, len(nums)) )
}


func max(x, y int) int {
    if x > y{return x}
    return y
}