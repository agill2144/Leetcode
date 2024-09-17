func findKthPositive(arr []int, k int) int {
    left := 0
    right := len(arr)-1
    offSet := 1
    for left <= right {
        mid := left + (right-left)/2
        correctIdx := arr[mid]-offSet
        diff := correctIdx-mid
        if diff < k {
            left = mid+1
        } else {
            right = mid-1
        }
    }
    rightVal := 1
    if right >= 0 {
        rightVal = arr[right]        
    }
    diff := (rightVal-offSet)-right
    return rightVal + (k-diff)
    
}