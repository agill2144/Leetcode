/**
 * // This is the MountainArray's API interface.
 * // You should not implement it, or speculate about its implementation
 * type MountainArray struct {
 * }
 *
 * func (this *MountainArray) get(index int) int {}
 * func (this *MountainArray) length() int {}
 */

func findInMountainArray(target int, mt *MountainArray) int {
    n := mt.length()
    left := 0
    right := n-1
    peakIdx := -1
    peakVal := -1
    for left <= right {
        mid := left + (right-left)/2
        midVal := mt.get(mid)
        if (mid == 0 || midVal > mt.get(mid-1)) && (mid == n-1 || midVal > mt.get(mid+1)) {
            peakIdx = mid
            peakVal = midVal
            break
        }
        if mid == n-1 || mt.get(mid+1) > midVal {
            left = mid+1
        } else {
            right = mid-1
        }

    }
    if target == peakVal {return peakIdx}

    left = 0
    right = peakIdx-1
    for left <= right {
        mid := left + (right-left)/2
        val := mt.get(mid)
        if val == target {return mid}
        if target > val {
            left = mid+1
        } else {
            right = mid-1
        }
    }

    // desc order side
    left = peakIdx
    right = n-1
    for left <= right {
        mid := left + (right-left)/2
        val := mt.get(mid)
        if val == target {return mid}
        if target > val {
            right = mid-1
        } else {
            left = mid+1
        }
    }
    return -1
}