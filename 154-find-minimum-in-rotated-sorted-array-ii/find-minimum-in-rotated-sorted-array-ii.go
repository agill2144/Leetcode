func findMin(nums []int) int {
    n := len(nums)
    left := 0
    right := n-1
    ans := math.MaxInt64
    for left <= right {
        mid := left + (right-left)/2
        if nums[mid] == nums[left] && nums[mid] == nums[right] {
            ans = min(ans, nums[mid])
            left++; right--
        }else if nums[left] <= nums[mid] {
            ans = min(ans, nums[left])
            left = mid+1
        } else {
            ans = min(ans, nums[mid])
            right = mid-1
        }
    }
    return ans
}
/*
    approach: linear scan
    - as soon as asc order breaks
    - we have found our min element
    time = o(n)
    space = o(1)
*/
// func findMin(nums []int) int {
//     for i := 1; i < len(nums); i++ {
//         if nums[i] < nums[i-1] {return nums[i]}
//     }
//     return nums[0]
// }