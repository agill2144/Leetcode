func findKthPositive(arr []int, k int) int {
    ptr := 1
    ms := 0
    i := 0
    for i < len(arr) {
        if arr[i] == ptr {
            i++; ptr++
        } else {
            ms++
            if ms == k {return ptr}
            ptr++
        }
    }
    rem := k-ms
    return ptr+rem-1
}