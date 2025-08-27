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
        // if we are looking at a sorted window
        // exit early because min will always be nums[left]
        if nums[mid] < nums[right] && nums[left] < nums[mid] {return nums[left]}

        // mid is min if mid < mid-1 and mid < mid+1
        if (mid == 0 || nums[mid] < nums[mid-1]) && (mid == n-1 || nums[mid] < nums[mid+1]) {
            return nums[mid]
        }
         

        if nums[left] <= nums[mid] {
            // left sorted, take search on right side
            left = mid+1
        } else {
            // right side is sorted, min is always on unsorted half
            // therefore take search on left side
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