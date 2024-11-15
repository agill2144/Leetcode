/*
    approach: binary search
    - mid is min if its value is smaller than its adj neighbors
    - min is always on the unsorted half
    time = o(logn)
    space = o(1)
*/
func findMin(nums []int) int {
    n := len(nums)
    left := 0
    right := n-1
    for left <= right {
        mid := left + (right-left)/2
        if nums[left] < nums[mid] && nums[mid] < nums[right] {return nums[left]}
        
        if (mid == 0 || nums[mid] < nums[mid-1]) && (mid == n-1 || nums[mid] < nums[mid+1]) {
            return nums[mid]
        } else if nums[left] <= nums[mid] {
            left = mid+1
        } else {
            right = mid-1
        }
    }
    return -1
}
/*
    approach: brute force
    - linear scan
    - until we find the first desc order
    - if in linear scan, we didnt find a desc order
    - then the whole array is sorted, therefore value at idx 0 is min
    time = o(n)
    space = o(1)
*/
// func findMin(nums []int) int {
//     for i := 1; i < len(nums); i++ {
//         if nums[i] < nums[i-1] {return nums[i]}
//     }
//     return nums[0]
// }