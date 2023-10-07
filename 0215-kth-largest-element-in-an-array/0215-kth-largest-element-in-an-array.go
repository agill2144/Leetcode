func findKthLargest(nums []int, k int) int {    
    left := 0
    right := len(nums)-1
    for left <= right {
        pivot := right
        nextSmaller := left
        for i := left; i < pivot; i++ {
            if nums[i] < nums[pivot] {
                nums[i], nums[nextSmaller] = nums[nextSmaller], nums[i]
                nextSmaller++
            }
        }
        nums[nextSmaller] , nums[pivot] = nums[pivot], nums[nextSmaller]
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