/*
    approach: binary search
    - numbers are sorted in the array
    - therefore something to do with binary search
    - and we also know the starting number is supposed to be 1
    - therefore we know that if we are at a idx ( mid ),
        - its correct value at that position is position is supposed to idx+1 ( i.e mid+1 )
    - NOW, we we need to see how many missing numbers we have had this far until mid
        - we have current val ( arr[mid] )
        - we have a "supposed-to-correct-val" (mid+1)
        - and now we need to check how many numbers are missing in-between them
        - that is, currentVal - correctVal
        - if this diff is less than k, it means that we do not have enough missing numbers
            - search right ; left = mid+1
        - otherewise, it means we have had more than k missing numbers on the left of mid
            - which means our kth is somewhere on the left hand side
            - search left ; right = mid-1
    - finally, when binary search breaks,
        - our answer will left + k
*/
func findKthPositive(arr []int, k int) int {
    left := 0
    right := len(arr)-1
    for left <= right {
        mid := left + (right-left)/2
        correctVal := mid+1
        currentVal := arr[mid]
        missingOnLeftOfMid := currentVal-correctVal
        
        if missingOnLeftOfMid < k {
            left = mid+1
        } else {
            right = mid-1
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