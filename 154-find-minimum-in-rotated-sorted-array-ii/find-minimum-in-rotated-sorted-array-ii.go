/*
    approach: binary search
    - following same pattern as "find min in rotated sorted array"
    - min element is always on the unsorted half ( when comparing sides with mid element )
    - but the input may contain duplicates
    - so we can do the pattern of finding element in rotated sorted array with duplicates
        - that is, if left == mid == right, 
            - move away from left and right ptrs and try again
            - BUT mid could be a min element ( so can left and so can right )
            - but left == mid and right == mid, therefore update our min so far with mid if needed
        - if left side is sorted
            - meaning left is <= mid element
            - min is going to be on the unsorted half
            - HOWEVER, because there are duplicates
            - left value could also be a min element
            - therefore before moving away from left , update our min answer if left is smaller
            - then take the search on right hand side
        - if right is sorted
            - meaning mid <= right element
            - min is always on unsorted half
            - but mid itself could be a min element
            - therefore before moving away from mid 
            - update our min ans if mid is new smaller element
            - could min be right?
            - no because mid is either < right or mid is == right
            - in the first case, we dont care about right, because right is bigger
            - in the second case, both elements are same, therefore we can again pick mid for min comparison
        
    tldr;
    - update min in all case
        1. when left == mid == right
        2. when left <= mid ( left could be a min element )
        3. when mid <= right ( mid could be a min element )
    time = o(logn)
    space = o(1)
*/
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