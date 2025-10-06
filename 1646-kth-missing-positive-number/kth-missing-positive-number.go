func findKthPositive(arr []int, k int) int {
    left := 0
    right := len(arr)-1
    if k < arr[0] {return k}
    for left <= right {
        mid := left + (right-left)/2
        correctVal := mid+1
        currentVal := arr[mid]
        miss := currentVal - correctVal
        if miss < k {
            left = mid+1
        } else {
            right = mid-1
        }
    }
    correctVal := right+1
    currentVal := arr[right]
    missing := currentVal - correctVal
    k -= missing
    return arr[right]+k
}