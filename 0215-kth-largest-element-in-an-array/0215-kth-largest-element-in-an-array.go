func findKthLargest(nums []int, k int) int {
    left := 0
    right := len(nums)-1
    for left <= right {
        ns := left
        pivotIdx := right
        for i := left; i < right; i++ {
            if nums[i] < nums[pivotIdx] {
                nums[ns], nums[i] = nums[i], nums[ns]
                ns++
            }
        }
        
        nums[ns], nums[pivotIdx] = nums[pivotIdx], nums[ns]
        if ns == len(nums)-k {
            return nums[ns]
        } else if len(nums)-k > ns {
            left = ns+1
        } else {
            right = ns-1
        }
    }
    return -1
}