func longestOnes(nums []int, k int) int {
    n := len(nums)
    if k >= n {return n}
    res := 0
    left := 0
    zeros := 0
    for i := 0; i < n; i++ {
        if nums[i] == 0 {
            zeros++
        }
        if zeros <= k {
            res = max(res, i-left+1)
        } else {
            if nums[left] == 0 {zeros--}
            left++
        }
    }
    return res
}