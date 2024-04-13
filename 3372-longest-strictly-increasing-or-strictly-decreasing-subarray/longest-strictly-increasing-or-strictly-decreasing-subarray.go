func longestMonotonicSubarray(nums []int) int {
    n := len(nums)
    inc := 1
    dec := 1
    ans := 1
    for i := 1; i < n; i++ {
        prev := nums[i-1]
        curr := nums[i]
        if curr > prev {
            inc++
            dec = 1
        } else if curr < prev {
            dec++
            inc = 1
        } else {
            inc = 1
            dec = 1
        }
        ans = max(ans, max(inc,dec))
    }
    return ans
}