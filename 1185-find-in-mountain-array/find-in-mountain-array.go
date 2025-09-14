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
        leftN := math.MinInt64
        rightN := math.MaxInt64
        if mid-1 >= 0 {leftN = mt.get(mid-1)}
        if mid+1 < n {rightN = mt.get(mid+1)}
        if midVal > leftN && midVal > rightN {
            peakIdx = mid
            peakVal = midVal
            break
        }
        if leftN > midVal {
            right = mid-1
        } else {
            left = mid+1
        }
    }
    if peakIdx == -1 {panic("dunno bruh")}
    if peakVal == target {return peakIdx}
    
    // target could be present on either side
    left = 0
    right = peakIdx-1
    for left <= right {
        mid := left + (right-left)/2
        midVal := mt.get(mid)
        if midVal == target {return mid}
        if target > midVal {
            left = mid+1
        } else {
            right = mid-1
        }
    }
    left = peakIdx+1
    right = n-1
    for left <= right {
        mid := left + (right-left)/2
        midVal := mt.get(mid)
        if midVal == target {return mid}
        if target > midVal {
            right = mid-1
        } else {
            left = mid+1
        }
    }
    return -1
}