func findKthLargest(nums []int, k int) int {
    left := 0
    right := len(nums)-1
    for left <= right {
        pivotIdx := right
        nextSmaller := left
        for i := left; i < pivotIdx; i++ {
            if nums[i] < nums[pivotIdx] {
                nums[i], nums[nextSmaller] = nums[nextSmaller], nums[i]
                nextSmaller++
            }
        }
        nums[pivotIdx], nums[nextSmaller] = nums[nextSmaller], nums[pivotIdx]
        if nextSmaller == len(nums)-k {
            return nums[nextSmaller]
        } else if len(nums)-k > nextSmaller {
            left = nextSmaller+1
        } else {
            right = nextSmaller-1
        }
    }
    return -1
}
