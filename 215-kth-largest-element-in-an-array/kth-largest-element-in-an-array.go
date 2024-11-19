func findKthLargest(nums []int, k int) int {
    n := len(nums)
    sort.Ints(nums)
    return nums[n-k]
}


/*
    k = 4

    1,2,2,3,3,4,5,5,6
*/