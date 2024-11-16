func findKthPositive(arr []int, k int) int {
    if k < arr[0] {return k}
    left := 0
    right := len(arr)-1
    for left <= right {
        mid := left + (right-left)/2
        cIdx := arr[mid]-1
        missing := cIdx - mid
        if missing >= k {
            right = mid-1
        } else {
            left = mid+1
        }
    }
    // we have already handled if missing number is to the left of the array
    // k < arr[0]
    // now this means, the missing number is either in the middle or to the right side of array
    missing := (arr[right]-1)-right
    rem := k-missing
    return arr[right]+rem
}



/*
    approach: linear scan with 2 ptrs
    time = o(n)
    space = o(1)
*/
// func findKthPositive(arr []int, k int) int {
//     ptr := 1
//     ms := 0
//     i := 0
//     for i < len(arr) {
//         if arr[i] == ptr {
//             i++; ptr++
//         } else {
//             ms++
//             if ms == k {return ptr}
//             ptr++
//         }
//     }
//     rem := k-ms
//     return ptr+rem-1
// }