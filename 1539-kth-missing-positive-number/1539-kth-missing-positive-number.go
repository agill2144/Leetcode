func findKthPositive(arr []int, k int) int {
    left := 0
    right := len(arr)-1
    for left <= right {
        mid := left + (right-left)/2
        val := arr[mid]
        correctIdx := val-1
        countMissingOnLeft := correctIdx - mid
        if countMissingOnLeft < k {
            left = mid+1
        } else {
            right = mid-1
        }
    }
    return left + k
}