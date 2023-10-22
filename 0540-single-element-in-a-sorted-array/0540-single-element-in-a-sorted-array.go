/*
    approach: brute force
    - linear scan
    time = o(n)
    space = o(1)
    
    approach: binary search ( sorted array )
    - The observation to notice is that there is a even,odd idx pattern
        - even,odd idx will have same repeating numbers
        - when this even,odd chain is broken by a single number, the chain becomes odd,even
        - therefore we need to find where our chain breaks the even,odd idx pattern
    - check if mid is our answer ?
        - mid will be our answer if it does not match its left and right adjacent neighbors
    - if mid is not our answer
    - if mid is at an odd idx
        - this means its prev element should be same value ( to follow even,odd idx pattern )
        - if prevVal == midVal , chain is not yet broken, search right
        - otherwise chain is broken, search left
    - if mid is at an even idx
        - this means its next element should be same value ( to follow even,odd idx )
        - if nextVal == midVal, chain is not yet broken, search right
        - otherwise chain is broken, search left
    
    time = o(logn)
    space = o(1)
*/
func singleNonDuplicate(nums []int) int {
    n := len(nums)
    left := 0
    right := n-1    
    for left <= right {
        mid := left + (right-left)/2
        midVal := nums[mid]
        nextVal := math.MaxInt64
        prevVal := math.MaxInt64
        
        // set next and prev value if we do not go out of bounds
        // if we do go out of bounds, its fine, there will be a mismatch with midVal
        // this happens when mid is at idx 0 or mid is at last idx
        // it means, we have not found the mismatch in between the array, 
        // the even-odd mismatch is lying on the edge of the array ( either idx 0 or last idx )
        if mid - 1 >= 0 {prevVal = nums[mid-1]}
        if mid + 1 < n  {nextVal = nums[mid+1]}
        
        if midVal != prevVal && midVal != nextVal {
            return midVal
        }
        
        if mid % 2 != 0 {
            if prevVal == midVal {
                left = mid+1
            } else {
                right = mid-1
            }
        } else {
            if nextVal == midVal {
                left = mid + 1
            } else {
                right = mid-1
            }
        }
    }
    return -1
}
