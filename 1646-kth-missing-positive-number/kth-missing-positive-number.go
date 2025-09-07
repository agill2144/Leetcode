func findKthPositive(arr []int, k int) int {
    n := len(arr)
    if k < arr[0] {return k}
    left := 0
    right := n-1
    for left <= right {
        mid := left+(right-left)/2
        correctVal := mid+1
        currentVal := arr[mid]
        missing := currentVal-correctVal
        if missing < k {
            left = mid+1
        } else {
            right = mid-1
        }
    }
    leftMissing := arr[right] - (right+1)
    k -= leftMissing
    return arr[right]+k
}