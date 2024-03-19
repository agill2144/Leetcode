func findKthPositive(arr []int, k int) int {
    mc := 0
    ptr := 1
    i := 0
    for i < len(arr) {
        if ptr != arr[i] {
            mc++
            if mc == k {
                return ptr+(k-mc)
            }
            ptr++
        } else {
            ptr++
            i++
        }
    }
    return ptr+(k-mc)-1
}