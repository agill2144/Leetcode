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

    // from right ptr, we want to figure out how many are missing on left of right
    // if right ptr went negative, it means num of missing on left of mid >= k
    // hence right ptr moved to left ( take binary search to left hand side )
    // however its possible right ptr goes negative
    // it means number of missing elements on left of right val should 0
    // since numbers start from 1, we can set right val to 1 ( if right ptr goes negative )
    // which will result into "count missing on left of right ptr = 0"
    rightVal := 1
    if right >=  0 {
        rightVal = arr[right]
    }
    missingOnLeft := (rightVal-1) - right
    // we are missing $missingOnLeft
    // remove those from k
    k -= missingOnLeft
    // now walk forward from right val 
    return rightVal + k
    
}