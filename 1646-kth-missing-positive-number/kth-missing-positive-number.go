func findKthPositive(arr []int, k int) int {
    if k < arr[0] {return k}
    left := 0
    right := len(arr)-1
    for left <= right {
        mid := left + (right-left)/2
        cIdx := arr[mid]-1
        miss := cIdx-mid
        if miss >= k {
            right = mid-1
        } else {
            left = mid+1
        }
    }
    missOnLeftOfRight := (arr[right]-1) - right
    updatedK := k - missOnLeftOfRight
    return arr[right] + updatedK
}

// k = 7
// [1,2,3,4]
// val = 5
// ptr = 4
// return val + (k-val)