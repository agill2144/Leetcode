/**
 * // This is the MountainArray's API interface.
 * // You should not implement it, or speculate about its implementation
 * type MountainArray struct {
 * }
 *
 * func (this *MountainArray) get(index int) int {}
 * func (this *MountainArray) length() int {}
 */

func findInMountainArray(target int, ma *MountainArray) int {
    n := ma.length()
    left := 0
    right := n-1
    peakIdx := -1
    for left <= right {
        mid := left + (right-left)/2
        midVal := ma.get(mid)
        if (mid == n || midVal > ma.get(mid+1)) && (mid == 0 || midVal > ma.get(mid-1)) {
            peakIdx = mid
            break
        } else if (mid == n-1 || ma.get(mid+1) > midVal) {
            left = mid+1
        } else {
            right = mid-1
        }
    }
    if peakIdx == -1 {return -1}

    // 0 -> peakIdx ( asc order )
    // peakIdx -> n-1 ( dsc order )
    left = 0
    right = peakIdx
    ans := -1
    for left <= right {
        mid := left + (right-left)/2
        midVal := ma.get(mid)
        if midVal == target {
            ans = mid
            right = mid-1
        } else if target > midVal {
            left = mid+1
        } else {
            right = mid-1
        }
    }
    if ans != -1 {return ans}

    left = peakIdx
    right = n-1
    for left <= right {
        mid := left + (right-left)/2
        midVal := ma.get(mid)
        if midVal == target {
            ans = mid
            right = mid-1
        } else if target > midVal {
            right = mid-1
        } else {
            left = mid+1
        }
    }
    return ans
}
