func rotate(nums []int, k int)  {
    k %= len(nums)
    
    // reverse whole arr
    nums = reverse(0, len(nums)-1, nums)
    
    // reverse 0:k-1 elements
    nums = reverse(0, k-1, nums)
    
    // reverse k:end elements
    nums = reverse(k, len(nums)-1, nums)
    
}


func reverse(left, right int, nums []int) []int {
    for left < right {
        nums[left], nums[right] = nums[right],nums[left]
        left++
        right--
    }
    return nums
}