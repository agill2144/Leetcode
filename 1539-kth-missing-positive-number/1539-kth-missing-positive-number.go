func findKthPositive(arr []int, k int) int {
    ptr := 1
    missingCount := 0
    i := 0
    for i < len(arr) {
        if arr[i] != ptr {
            missingCount++
            if missingCount == k {
                return ptr
            }
            // check the next value
            ptr++
        } else {
            ptr++
            i++
        }
    }
    return arr[len(arr)-1] + (k-missingCount)  
}