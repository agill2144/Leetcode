func rotate(nums []int, k int)  {
    n := len(nums)
    if k > n {k = k % n}
    // reverse full
    reverse(0, n-1,nums)
    // reverse first k
    reverse(0,k-1,nums)
    // reverse remaining after first k elements
    reverse(k, n-1,nums)
}

func reverse(left, right int, nums []int) {
    for left < right {
        nums[left], nums[right] = nums[right], nums[left]
        left++
        right--
    }
}