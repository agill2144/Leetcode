func rotate(nums []int, k int)  {
    if k >= len(nums) { k = k % len(nums) }
    
    // reverse whole array
    reverse(0, len(nums)-1,nums)

    // reverse first k elements
    reverse(0, k-1,nums)
    
    // reverse k: elements
    reverse(k, len(nums)-1,nums)
}

func reverse(left, right int, nums []int) {
    for left <= right {
        nums[left], nums[right] = nums[right], nums[left]
        left++
        right--
    }
}