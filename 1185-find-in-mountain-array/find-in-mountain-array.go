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
        val := mt.get(mid)
        prev := math.MinInt64
        next := math.MinInt64
        if mid-1 >= 0 {prev = mt.get(mid-1)}
        if mid+1 < n {next = mt.get(mid+1)}
        if val > prev && val > next {
            peakIdx = mid
            peakVal = val
            break
        }
        if prev > val {
            right = mid-1
        } else {
            left = mid+1
        }
    }
    if peakIdx == -1 {panic("well shit")}
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