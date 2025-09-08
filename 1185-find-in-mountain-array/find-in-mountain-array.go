/**
 * // This is the MountainArray's API interface.
 * // You should not implement it, or speculate about its implementation
 * type MountainArray struct {
 * }
 *
 * func (this *MountainArray) get(index int) int {}
 * func (this *MountainArray) length() int {}
 */

func findInMountainArray(target int, m *MountainArray) int {
    // 1. find peak, so that we have two sorted arrays
    // - left side is increasing
    // - right side is decreasing
    // could there be more than 1 peak? 
    // - not clear from the question, but examples suggest that 
    // - there wont be more than 1 peak
    // mid is peak if its greater than both its adj neis
    n := m.length()
    left := 0
    right := n-1
    peakIdx := -1
    peakVal := math.MinInt64
    for left <= right {
        mid := left + (right-left)/2
        midVal := m.get(mid)
        if (mid == 0 || midVal > m.get(mid-1)) && (mid == n-1 || midVal > m.get(mid+1)) {
            peakIdx = mid
            peakVal = midVal
            break
        }
        if mid == n-1 || m.get(mid+1) > midVal {
            left = mid+1
        } else {
            right = mid-1
        }

    }
    if peakIdx == -1 {return -1}
    if peakVal == target {return peakIdx}
    // search in left half, which is sorted in asc order
    left = 0
    right = peakIdx-1
    for left <= right {
        mid := left + (right-left)/2
        midVal := m.get(mid)
        if midVal == target {
            return mid
        } else if target > midVal {
            left = mid+1
        } else {
            right = mid-1
        }
    }

    // search in right half, right side is sorted in desc order
    left = peakIdx+1
    right = n-1
    for left <= right {
        mid := left + (right-left)/2
        midVal := m.get(mid)
        if midVal == target {
            return mid
        } else if target > midVal {
            right = mid-1
        } else {
            left = mid+1
        }
    }
    return -1
}