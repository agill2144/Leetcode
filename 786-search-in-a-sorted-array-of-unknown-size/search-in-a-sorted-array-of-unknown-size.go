/**
 * // This is the ArrayReader's API interface.
 * // You should not implement it, or speculate about its implementation
 * type ArrayReader struct {
 * }
 *
 * func (this *ArrayReader) get(index int) int {}
 */

func search(reader ArrayReader, target int) int {
    left := 0
    right := 1
    for reader.get(right) < target {
        right *= 2
    }
    for left <= right {
        mid := left + (right-left)/2
        val := reader.get(mid)
        if val == target {return mid}
        if val < target {
            left = mid+1
        } else {
            right = mid-1
        }
    }
    return -1
}