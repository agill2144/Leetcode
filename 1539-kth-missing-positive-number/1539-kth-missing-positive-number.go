func findKthPositive(arr []int, k int) int {
    left := 0
    right := len(arr)-1
    for left <= right {
        mid := left + (right-left)/2
        correctVal := mid+1
        currentVal := arr[mid]
        missingOnLeftOfMid := currentVal-correctVal

        if missingOnLeftOfMid >= k {
            // number of missing numbers are more than k,
            // therefore our kth missing will be on left side
            right = mid-1
        } else {
            // number of missing numbers are less than k,
            // therefore we need to search in an area where there are >= k missing numbers
            // therefore search on right side
            left = mid+1
        }
    }
    return left + k
}
/*
    approach: brute force; linear scan
    - missing number could be missing on the left side, in-between or on the right side of the array
    - we know that the list is supposed to start with ele = 1
    - and we know the list is also sorted
    - with missing numbers in between ( maybe no missing nums in between )
    - we can use a ptr that acts as the desired number starting from 1
    - and use another ptr to check if desiredNumber matches with whats in the arr
    - if the desiredNumber != whats in the arr, it means we have a missing number
        - therefore increment missing counter by 1
        - if our missing counter equals k, we have found the missing number ( i.e the current desired number )
        - otherwise, move our desiredNumber ptr forward
    - if the desiredNumber does match with whats in arr, we can move desiredNumber and i ptr forward.
    
    time = o(maxNumber)
    space = o(1)
*/
// func findKthPositive(arr []int, k int) int {
//     ptr := 1
//     missingCount := 0
//     i := 0
//     for i < len(arr) {
//         if arr[i] != ptr {
//             missingCount++
//             if missingCount == k {
//                 return ptr
//             }
//             // check the next value
//             ptr++
//         } else {
//             ptr++
//             i++
//         }
//     }
//     return arr[len(arr)-1] + (k-missingCount)  
// }