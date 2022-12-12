// time: o(logn)
// space: o(1)
func fixedPoint(arr []int) int {
    if arr == nil || len(arr) == 0 {return -1}
    left := 0
    right := len(arr)-1
    idx := -1
    for left <= right {
        mid := left + (right-left) / 2
        if arr[mid] == mid {
            idx = mid
            right = mid-1
        /*
            if the value at mid is less than mid idx,
            then and if we go left more, the mid idx will reduce, but the value will also decrease (because its sorted)
            so if the value at mid idx is < idx, going left more will never make them meet because the value is already
            less than current mid position. Why? because arr is sorted and does not have dupes.
        */
        } else if arr[mid] < mid {
            left = mid+1
        } else {
            right = mid-1
        }
    }
    return idx
}