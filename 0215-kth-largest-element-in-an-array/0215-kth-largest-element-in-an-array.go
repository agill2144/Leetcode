
/*
    brute force:
    - sort the array
    - return len(arr)-k idx element
    time = o(nlogn)
    space = maybe o(n) if sort is used mergeSort alg
    
    but what is k is larger than the input array
    - make k relative to the size of arry ; k = k % len(arr)
*/
func findKthLargest(nums []int, k int) int {
    if nums == nil || len(nums) == 0 {return -1}
    if len(nums) == 1 {return nums[0]}
    sort.Ints(nums)
    return nums[len(nums)-k]
}