func longestMonotonicSubarray(nums []int) int {
    n := len(nums)
    res := 1
    incr := 1
    decr := 1
    for i := 1; i < n; i++ {
        if nums[i] > nums[i-1] { incr++ } else { incr = 1 }
        if nums[i] < nums[i-1] { decr++ } else { decr = 1 }
        res = max(res, max(incr, decr))
    }
    return res
}