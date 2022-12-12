/*
    approach: brute force
    - linear scan, left to right
    - if num != prev/next number, then this is the single occurence number
    time: o(n)
    space: o(1)
    
    approach: binary search
    - the numbers are in sorted, so somehting to do with binary search to optimize linear o(n) solution
    - If you notice, if all of the numbers were duped, then each number would take even and odd index - in this format
    - nums = [1,1,2,2,3,3,4,4]
    - index = 0,1,2,3,4,5,6,7
    - And when a number only shows up once and rest show up twice
    - nums = [1,1,2,2,3,4,4]
    - index = 0,1,2,3,4,5,6
    - then the format becomes even->odd <sing-element> odd->even
    - Thats what we need to find with binary search, where does the pattern from even odd index changes to odd->even
    - How do we know if mid is our answer?
        - if value at mid is not same as previous and next, then mid value is the single occurring element
    - How do we decide to search left or right?
        - if mid idx happens to land on an even idx and mid+1 idx is the same value as mid idx, or
        - or if mid idx happens to land on an odd idx and mid-1 (its even pair) is the same value as mid idx, 
            - that means so far the pattern has not flipped, search right side of mid
            - why? because when that pattern of even-odd breaks, the pattern becomes odd-even
            - how do we check that?
            - if mid idx is even, that means, the next element SHOULD BE THE SAME ( because the natural order of dupes is even-odd)
            - if mid idx is odd, that means, the previous element SHOULD BE THE SAME 
            - if the above 2 are true, that means a single element has not flipped the even->odd pattern - so search right
        - otherwise search left ( odd->even detected )
    
    time: o(logn)
    space: o(1)
*/
func singleNonDuplicate(nums []int) int {
    left := 0
    right := len(nums)-1
    
    for left <= right {
        mid := left + (right-left) / 2
        if (mid==0 ||nums[mid] != nums[mid-1]) && (mid==len(nums)-1||nums[mid] != nums[mid+1]) {
            return nums[mid]
        } else if (mid % 2 == 0 && (mid==len(nums)-1||nums[mid+1] == nums[mid])) || (mid % 2 != 0 && (mid==0||nums[mid] == nums[mid-1])) {
            left = mid+1
        } else {
            right = mid-1
        }
    }
    return -1
}