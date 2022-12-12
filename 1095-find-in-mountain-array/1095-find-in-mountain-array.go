/**
 * // This is the MountainArray's API interface.
 * // You should not implement it, or speculate about its implementation
 * type MountainArray struct {
 * }
 *
 * func (this *MountainArray) get(index int) int {}
 * func (this *MountainArray) length() int {}
 */

func findInMountainArray(target int, mountainArr *MountainArray) int {
    n := mountainArr.length()
    left := 0
    right := n-1
    peakIdx := -1
    peakVal := -1
    for left <= right {
        mid := left+(right-left)/2
        midVal := mountainArr.get(mid)
        if (mid==0|| midVal > mountainArr.get(mid-1)) && (mid==n-1|| midVal > mountainArr.get(mid+1)) {
            peakIdx = mid
            peakVal = midVal
            break
        } else if (mid==0|| midVal > mountainArr.get(mid-1)) {
            left = mid+1
        } else { 
            right = mid-1
        }
    }
    if peakIdx == -1 {return -1}
    if peakVal == target {return peakIdx}
    
    // search left side of the mountain ( asc order )
    left = 0
    right = peakIdx
    for left <= right {
        mid := left + (right-left)/2
        midVal := mountainArr.get(mid)
        if midVal == target {return mid}
        if target > midVal {
            left = mid+1
        } else {
            right = mid-1
        }
    }
    
    // if we did not exit from left side search, then search right side of the mountain ( desc order )
    left = peakIdx
    right = n-1
    for left <= right {
        mid := left + (right-left)/2
        midVal := mountainArr.get(mid)
        if midVal == target {return mid}
        if target < midVal {
            left = mid+1
        } else {
            right = mid-1
        }
    }
    return -1
}