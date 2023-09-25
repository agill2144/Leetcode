func rotate(nums []int, k int)  {
    // what if k > len(nums)
    // - ensure k is relative to the size of the arr
    if k > len(nums) {k = k % len(nums)}
    // reverse entire array first 
    // so that the last k are now at the front of the arr
    reverse(0, len(nums)-1, nums)
    // reverse first k ( idx 0 to k-1 )
    reverse(0, k-1, nums)
    // undo the first reverse from kth idx:end
    reverse(k, len(nums)-1, nums)
    
}

func reverse(left, right int, nums []int) {
    for left < right {
        nums[left], nums[right] = nums[right], nums[left]
        left++
        right--
    }
} 