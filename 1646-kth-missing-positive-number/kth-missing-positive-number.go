func findKthPositive(arr []int, k int) int {
    ptr := 1
    ptr2 := 0
    count := 0
    for ptr2 < len(arr) {
        if ptr != arr[ptr2] {
            count++
            if count == k {return ptr}
            ptr++
        } else {
            ptr2++
            ptr++
        }
    }
    return arr[len(arr)-1] + (k-count)
}