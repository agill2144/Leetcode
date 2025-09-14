func findKthPositive(arr []int, k int) int {
    if k < arr[0] {return k}
    n := len(arr)
    left := 0
    right := n-1
    for left <= right {
        mid := left + (right-left)/2
        ctv := mid+1
        cv := arr[mid]
        miss := cv - ctv
        if miss >= k {
            right = mid-1
        } else {
            left = mid+1
        }
    }
    ctv := right+1
    cv := arr[right]
    miss := cv-ctv
    updatedK := k - miss
    return arr[right] + updatedK
}