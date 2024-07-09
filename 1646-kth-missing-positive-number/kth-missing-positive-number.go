func findKthPositive(arr []int, k int) int {
    // [2,3,4,7,8]
    ptr := 1
    missingCount := 0
    i := 0
    for i < len(arr) {
        ithVal := arr[i]
        if ithVal == ptr {
            i++
            ptr++
        } else {
            missingCount++
            if missingCount == k {
                return ptr
            }
            ptr++
        }
    }
    remainCount := k-missingCount
    return arr[len(arr)-1]+remainCount
}