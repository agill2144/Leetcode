func findKthPositive(arr []int, k int) int {
    left := 0
    right := len(arr)-1
    for left <= right {
        mid := left + (right-left)/2
        correct := mid+1
        current := arr[mid]
        missingOnLeft := current-correct
        if missingOnLeft < k {
            left = mid+1
        } else {
            right = mid-1
        }
    }
    return left+k
}
// time = max(n,k)
// space = o(1)
// func findKthPositive(arr []int, k int) int {
//     // [100] k = 200
//     ptr := 1
//     missingCount := 0
//     i := 0
//     for i < len(arr) {
//         ithVal := arr[i]
//         if ithVal == ptr {
//             i++
//             ptr++
//         } else {
//             missingCount++
//             if missingCount == k {
//                 return ptr
//             }
//             ptr++
//         }
//     }
//     remainCount := k-missingCount
//     return arr[len(arr)-1]+remainCount
// }