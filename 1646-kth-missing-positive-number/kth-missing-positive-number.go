func findKthPositive(arr []int, k int) int {
    n := len(arr)
    left := 0
    right := n-1
    if k < arr[0] {return k}
    for left <= right {
        mid := left + (right-left)/2
        correctVal := mid+1
        currentVal := arr[mid]
        missingOnLeft := currentVal-correctVal
        if missingOnLeft >= k {
            right = mid-1
        } else {
            left = mid+1
        }
    }
    correctVal := right+1
    currentVal := arr[right]
    missing := currentVal-correctVal
    k -= missing
    return arr[right]+k
}