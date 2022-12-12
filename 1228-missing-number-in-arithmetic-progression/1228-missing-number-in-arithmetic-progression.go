func missingNumber(arr []int) int {
    n := len(arr)
    diff := (arr[n-1]-arr[0])/n
    left := 0
    right := n-1
    
    
    for left <= right {
        mid := left+(right-left)/2
        if arr[mid] == arr[0]+(mid*diff) {
            left = mid+1
        } else {
            right = mid-1
        }
    }
    return arr[right]+diff
}

