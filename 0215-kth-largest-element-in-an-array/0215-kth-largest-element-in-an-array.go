func findKthLargest(nums []int, k int) int {
    left := 0
    right := len(nums)-1
    
    for left <= right {
        
        pivotIdx := right
        ns := left
        for i := left; i < pivotIdx; i++ {
            if nums[i] < nums[pivotIdx] {
                nums[i], nums[ns] = nums[ns], nums[i]
                ns++
            }
        }
        nums[ns], nums[pivotIdx] = nums[pivotIdx], nums[ns]
        if ns == len(nums)-k {return nums[ns]}
        if len(nums)-k > ns {
            left = ns+1
        } else {
            right = ns-1
        }
    }
    return -1
}