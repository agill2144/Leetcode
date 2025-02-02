func findKthPositive(arr []int, k int) int {
    if k < arr[0] {return k}
    left := 0
    right := len(arr)-1
    for left <= right {
        mid := left + (right-left)/2
        missOnLeft := (arr[mid]-1)-mid
        if k <= missOnLeft {
            right = mid-1
        } else {
            left = mid+1
        }
    }
    missOnLeft := (arr[right]-1)-right
    k -= missOnLeft
    return arr[right]+k
}