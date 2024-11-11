func rotate(nums []int, k int)  {
    n := len(nums)
    k %= n
    // k = 3
    // [1,2,3,4,5,6,7]
    // [7,6,5,4,3,2,1]
    reverse(nums, 0, n-1)
    // [5,6,7,4,3,2,1]
    reverse(nums, 0, k-1)
    reverse(nums, k, n-1)
    
}

func reverse(nums []int, left, right int) {
    for left < right {
        nums[left], nums[right] = nums[right], nums[left]
        left++; right--
    }
}