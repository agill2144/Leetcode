func minOperations(nums []int, x int) int {
    if nums[0] > x && nums[len(nums)-1] > x {return -1}
    total := 0
    for i := 0; i < len(nums); i++ {total += nums[i]}
    k := total - x
    left := 0
    sum := 0
    ans := len(nums)+1
    for i := 0; i < len(nums); i++ {
        sum += nums[i]
        for left <= i && sum > k {
            sum -= nums[left]
            left++
        }
        if sum == k {
            ans = min(ans, len(nums)-(i-left+1))
        }
    }
    if ans == len(nums)+1 {return -1}
    return ans
}