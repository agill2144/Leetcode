/*
    approach: binary search
    - put the input on a board / paper
    - mark the indicies down for each element
    - try to observe / look for a pattern

    - the pattern is even/odd index
    - that is, every element at even idx has its dupe pair at next odd index
    - this pattern follows even/odd index until the interruption happens
    - which means, if at mid the pattern is interrupted ( even/odd pattern )
    - it means, the interruption has had happened already, look left
    - otherwise the interruption has not yet happened therefore even/odd pattern is not yet broken
        - therefore continue searching right

    - how do we know if val mid is our ans tho ?
    - if mid does not equal any of its immediate ( left, right ) neighbors,
        - it means its the only single element
    
    The input could have been unsorted btw, we can always sort
    - this question may get twisted like this ^

    time = o(logn)
    space = o(1)
*/
func singleNonDuplicate(nums []int) int {
    n := len(nums)
    left := 0
    right := n-1
    for left <= right {
        mid := left + (right-left)/2

        if (mid == 0 || nums[mid] != nums[mid-1]) && (mid == n-1 || nums[mid] != nums[mid+1]) {
            return nums[mid]
        }

        if mid % 2 == 0 {
            if (mid == n-1 || nums[mid] != nums[mid+1]) {
                right = mid-1
            } else {
                left = mid+1
            }
        } else if mid % 2 != 0 {
            if (mid == 0 || nums[mid] != nums[mid-1]) {
                right = mid-1
            } else {
                left = mid+1
            }
        }
    }
    return -1
}