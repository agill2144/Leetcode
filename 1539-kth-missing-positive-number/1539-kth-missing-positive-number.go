func findKthPositive(arr []int, k int) int {
    left := 0
    right := len(arr)-1
    for left <= right {
        mid := left + (right-left)/2
        val := arr[mid]
        desiredIdx := val-1
        countMissingOnLeft := desiredIdx - mid
        if countMissingOnLeft < k {
            // there are less than k numbers missing to the left of mid,
            left = mid+1
        } else {
            // there are more than k numbers missing to the left of mid,
            // we want kth missing, therefore search left
            right = mid-1
        }
    }
    return left + k
}