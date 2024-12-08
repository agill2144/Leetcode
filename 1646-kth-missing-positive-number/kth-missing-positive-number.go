func findKthPositive(arr []int, k int) int {
    if k < arr[0] {
        return k
    }
    left := 0
    right := len(arr)-1
    for left <= right {
        mid := left + (right-left)/2
        correctIdx := arr[mid]-1
        missingOnLeft := correctIdx - mid
        if k <= missingOnLeft {
            right = mid-1
        } else {
            left = mid+1
        }
        //  0 1 2 3 4
        // [2,3,4,7,11]
        // [1,2,3,4,5]
    }

    correctRightIdx := arr[right]-1
    missingOnLeft := correctRightIdx - right
    k -= missingOnLeft
    return arr[right]+k
}