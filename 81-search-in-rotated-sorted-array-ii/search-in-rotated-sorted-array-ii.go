/*
    approach: brute force
    - linear scan ; left to right
    time = o(n)
    space = o(1)

    Because the array was at some point sorted
    and then rotated, we must try binary search

    approach: binary search
    - if mid is our ans, return true
    - look for sorted half from mid's perspective
    - i.e is left half sorted or right half sorted
    - Note: there are dupes; so left could be <= mid and/or right >= mid
    - [8 . . . . . 8 . . . .  2]
       l           m          r
    - in the above case, left side is sorted, from mids prespective

    - [8 . . . . . 2 . . . .  2]
       l           m          r
    - in the above case, right side is sorted, from mids perspective
    - then check whether our target is in the sorted half
        - if yes, take the search to that side
        - if no,  take the search to other side
    
    - what happens when we cannot determine FOR SURE a side is sorted
    - i.e left == mid == right
    - [8 . . . . . 8 . . 2 .  8]
       l           m          r
    - We cannot confidently say search in the left half or the right half
    - when this happens, shrink our window down from both sides , in attempt to get out of this situation
    - left++; right--, and re-calc mid idx
    - [8 8 . . . . 8 . . 2 2 8]
         l         m       r
    - Now we can say 100% say that left side is sorted from mid's perspective
    - check if the target is in sorted half and take the search to either side accordingly
    
    ** ALWAYS USE MID TO FIGURE OUT WHICH SIDE TO TAKE THE SEARCH ON **
    Sometimes, mid's immediate neighbors are used
    Sometimes, only mid is used
    Sometimes, idx positions are used
    Sometimes, left and right ptrs are used
    - But in all cases, its compared/checked from mids perspective

    time = o(logn) for best case
            o(n) for worst case [1,1,1,1,1,1,1] ; target = 2
    space = o(1)
*/
func search(nums []int, target int) bool {
    left := 0
    right := len(nums)-1
    for left <= right {
        mid := left + (right-left)/2
        if nums[mid] == target {return true}
        
        if nums[mid] == nums[left] && nums[mid] == nums[right] {
            left++; right--; continue
        }
        
        if nums[left] <= nums[mid] {
            if target <= nums[mid] && target >= nums[left] {
                right=mid-1
            } else {
                left = mid+1
            }
        } else {
            if target >= nums[mid] && target <= nums[right] {
                left = mid+1
            } else {
                right = mid-1
            }
        }
    }
    return false
}